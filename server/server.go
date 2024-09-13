package server

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"log/slog"
	"slices"

	"github.com/scatternoodle/wflang/internal/jrpc2"
	"github.com/scatternoodle/wflang/internal/lsp"
	"github.com/scatternoodle/wflang/wflang/ast"
	"github.com/scatternoodle/wflang/wflang/parser"
	"github.com/scatternoodle/wflang/wflang/token"
)

var debug bool

func New(name, version *string, dbg bool) *Server {
	debug = dbg

	srv := &Server{
		name:         name,
		version:      version,
		trace:        lsp.TraceOff,
		uri:          "",
		initialized:  false,
		capabilities: serverCapabilities(),
		parser:       nil,
		ast:          nil,
	}

	srv.handlers = map[string]handlerFunc{
		lsp.MethodInitialize:         srv.handleInitializeRequest,
		lsp.MethodInitialized:        srv.handleInitializedNotification,
		lsp.MethodDocDidOpen:         srv.handleDocDidOpenNotification,
		lsp.MethodDocDidChange:       srv.handleDocDidChangeNotification,
		lsp.MethodDocDidSave:         srv.handleDocDidSaveNotification,
		lsp.MethodSemanticTokensFull: srv.handleSemanticTokensFullRequest,
		lsp.MethodHover:              srv.handleHoverRequest,
		lsp.MethodShutdown:           srv.handleShutdownRequest,
		lsp.MethodExit:               srv.handleExitNotification,
		lsp.MethodDocumentSymbols:    srv.handleDocumentSymbolsRequest,
		lsp.MethodDefinition:         srv.handleGotoDefinitionRequest,
		lsp.MethodCompletion:         srv.handleCompletionRequest,
		lsp.MethodRename:             srv.handleRenameRequest,
		lsp.MethodSignatureHelp:      srv.handleSignatureHelpRequest,
		lsp.MethodSetTrace:           srv.handleSetTraceNotification,
	}
	return srv
}

type Server struct {
	name    *string
	version *string
	trace   lsp.TraceValue
	// for now, server only handles a single document - this likely will need to turn into a map[string]*parser.Parser at some point
	uri          string
	capabilities lsp.ServerCapabilities
	initialized  bool // before this is set true, we only accept requests with initialize method
	exiting      bool // set after an shutdown request is received, awaiting exit request
	parser       *parser.Parser
	ast          *ast.AST
	handlers     map[string]handlerFunc
	symbols      map[string]lsp.DocumentSymbol

	*tokenEncoder
}

func serverCapabilities() lsp.ServerCapabilities {
	return lsp.ServerCapabilities{
		TextDocumentSync: lsp.SyncFull,
		SemanticTokensProvider: lsp.SemanticTokensOptions{
			Legend: lsp.TokenTypesLegend{
				TokenTypes:     tokenTypes(),
				TokenModifiers: tokenModifiers(),
			},
			Range: false,
			Full:  true,
		},
		HoverProvider:          true,
		DocumentSymbolProvider: true,
		DefinitionProvider:     true,
		CompletionProvider: lsp.CompletionOptions{
			CompletionItem: &lsp.CompletionItemOptions{LabelDetailsSupport: true},
			// TODO we'll want trigger on at least ',' once methods implemented
		},
		RenameProvider: true,
		SignatureHelpProvider: &lsp.SignatureHelpOptions{
			TriggerChars:   []string{"("},
			RetriggerChars: nil,
		},
	}
}

func (srv *Server) setInit(req lsp.InitializeRequestParams) error {
	return srv.setTrace(req.Trace)
}

func (srv *Server) setTrace(t lsp.TraceValue) error {
	switch t {
	case lsp.TraceOff, lsp.TraceMessages, lsp.TraceVerbose:
		slog.Debug("setting trace", "before", srv.trace, "after", t)
		srv.trace = t
	default:
		return fmt.Errorf("unrecognized trace value: %s", t)
	}
	return nil
}

func (srv *Server) initializeResponse(id *int) lsp.InitializeResponse {
	var srvInfo *lsp.AppInfo
	if srv.name == nil {
		srvInfo = nil
	} else {
		ver := ""
		if srv.version != nil {
			ver = *srv.version
		}
		srvInfo = &lsp.AppInfo{Name: *srv.name, Version: ver}
	}

	return lsp.InitializeResponse{
		Response: jrpc2.NewResponse(id, nil),
		Result: lsp.InitializeResult{
			Capabilities: srv.capabilities,
			ServerInfo:   srvInfo,
		},
	}
}

func (srv *Server) ListenAndServe(r io.Reader, w io.Writer) {
	slog.Info("Scanning for messages...")
	scanner := bufio.NewScanner(r)
	scanner.Split(jrpc2.Split)

	for scanner.Scan() {
		srv.handleMessage(w, scanner.Bytes())
	}
	slog.Info("Server stopped listening")
}

func (srv *Server) handleMessage(w io.Writer, msg []byte) {
	method, content, err := jrpc2.DecodeMessage(msg)
	if err != nil {
		slog.Error("Unable to decode", "error", err, "method", method, "message", msg)
		return
	}

	requestId := getRequestID(content)
	slog.Info("Recieved", "method", method, "id", func() string {
		if requestId == nil {
			return ""
		}
		return fmt.Sprint(*requestId)
	}())
	slog.Debug(fmt.Sprintf("Content=%s", string(content)))

	if !srv.initialized && method != lsp.MethodInitialize && method != lsp.MethodInitialized {
		respondError(w, requestId, lsp.ERRCODE_SERVER_NOT_INITIALIZED, "server not yet initialized", nil)
		return
	}

	if srv.exiting && method != lsp.MethodShutdown && method != lsp.MethodExit {
		respondError(w, requestId, jrpc2.ERRCODE_INVALID_REQUEST, "server is shutting down, expects 'exit' method", nil)
		return
	}

	handler, ok := srv.handlers[method]
	if !ok {
		slog.Warn("Unhandled method", "method", method, "id", requestId)
		return
	}
	handler(w, content, requestId)
}

func (srv *Server) getTokenAtPos(pos lsp.Position) (index int, tok token.Token, ok bool) {
	toks := srv.parser.Tokens()
	idx := slices.IndexFunc(toks, func(t token.Token) bool {
		if t.StartPos.Line != pos.Line {
			return false
		}
		return pos.Character >= t.StartPos.Col && pos.Character <= t.EndPos.Col
	})

	if idx < 0 {
		return -1, token.Token{}, false
	}
	return idx, toks[idx], true
}

func send(w io.Writer, v any) {
	response, err := jrpc2.EncodeMessage(v)
	if err != nil {
		panic(fmt.Errorf("response failed during encoding: %w", err))
	}
	if _, err = w.Write(response); err != nil {
		panic(fmt.Errorf("response failed during write: %w", err))
	}
	slog.Debug(fmt.Sprintf("Wrote content=%s", string(response)))
}

func respondError(w io.Writer, id *int, code int, msg string, dat ...any) {
	rErr := jrpc2.ResponseError{
		Code:    code,
		Message: msg,
		Data:    dat,
	}
	v := jrpc2.NewResponse(id, &rErr)
	send(w, v)
}

// getRequestID returns the Request ID from a content byte slice. Returns null if
// unable to resolve the Request ID.
func getRequestID(b []byte) *int {
	var idObj struct {
		ID *int `json:"id,omitempty"` // We want to be pretty permissive here as we can just return null if we can't find
	} // the ID.
	if err := json.Unmarshal(b, &idObj); err != nil {
		return nil
	}
	return idObj.ID
}

// Unmarshals an LSP request/notification message into the given interface. Returns
// true if successful, else responds to the message with an error before returning
// false.
func handleParseContent(v any, w io.Writer, c []byte, id *int) bool {
	if err := json.Unmarshal(c, v); err != nil {
		var idStr string
		if id == nil {
			idStr = "nil"
		} else {
			idStr = fmt.Sprint(*id)
		}
		slog.Error("error parsing request", "id", idStr, "err", err)
		respondError(w, id, lsp.ERRCODE_REQUEST_FAILED, fmt.Sprintf("parse error: %s", err), nil)
		return false
	}

	return true
}

// handleAssertID returns true if the given id is non-nil, or else response with
// the appropriate error and returns false.
func handleAssertID(w io.Writer, id *int) bool {
	if id == nil {
		respondError(w, id, lsp.ERRCODE_REQUEST_FAILED, "request ID cannot be nil", nil)
		return false
	}
	return true
}

// cursorPos returns the "cursor position" of a given lsp.Position, which is
// at column-1. If column is zero, returns zero.
func cursorPos(pos lsp.Position) lsp.Position {
	cursorPos := pos
	if cursorPos.Character > 0 {
		cursorPos.Character--
	}
	return cursorPos
}

// logTrace creates and sends an lsp.LogTraceNotification on the given writer.
// Message will always be sent, wherease the verbose param is only send it the
// server's trace setting is on "verbose".
func (s *Server) logTrace(w io.Writer, message, verbose string) {
	if s.trace == lsp.TraceOff {
		return
	}
	trace := lsp.LogTraceNotification{
		Notification: jrpc2.NewNotification(lsp.MethodLogTrace),
		LogTraceParams: lsp.LogTraceParams{
			Message: message,
			Verbose: verbose,
		},
	}
	send(w, trace)
}
