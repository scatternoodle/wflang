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
	"github.com/scatternoodle/wflang/lang/parser"
	"github.com/scatternoodle/wflang/lang/token"
)

var debug bool

func New(name, version *string, debug bool) *Server {
	srv := &Server{
		name:         name,
		version:      version,
		initialized:  false,
		capabilities: serverCapabilities(),
		parser:       nil,
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
	}
	return srv
}

type Server struct {
	name         *string
	version      *string
	capabilities lsp.ServerCapabilities
	initialized  bool // before this is set true, we only accept requests with initialize method
	exiting      bool // set after an shutdown request is received, awaiting exit request
	parser       *parser.Parser
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
		CompletionProvider:     lsp.CompletionOptions{
			// TODO we'll want trigger on at least ',' once methods implemented
		},
	}
}

func (srv *Server) initialize(id *int) lsp.InitializeResponse {
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
		debugNotification(w, fmt.Sprintf("unhandled method: %s", method))
		return
	}
	handler(w, content, requestId)
}

func (srv *Server) getTokenAtPos(pos lsp.Position) (tok token.Token, ok bool) {
	toks := srv.parser.Tokens()
	idx := slices.IndexFunc(toks, func(t token.Token) bool {
		if t.StartPos.Line != pos.Line {
			return false
		}
		return pos.Character >= t.StartPos.Col && pos.Character <= t.EndPos.Col
	})

	if idx < 0 {
		return token.Token{}, false
	}
	return toks[idx], true
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

func respondError(w io.Writer, id *int, code int, msg string, dat any) {
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
		respondError(w, id, jrpc2.ERRCODE_PARSE_ERROR, "parse error", err)
		return false
	}

	return true
}

// handleAssertID returns true if the given id is non-nil, or else response with
// the appropriate error and returns false.
func handleAssertID(w io.Writer, id *int) bool {
	if id == nil {
		respondError(w, id, jrpc2.ERRCODE_INVALID_REQUEST, "request ID cannot be nil", nil)
		return false
	}
	return true
}

// debugNotification creates and sends an lsp.ShowMessageNotification on w with
// message msg.
func debugNotification(w io.Writer, msg string) {
	if !debug {
		return
	}

	not := lsp.ShowMessageNotification{
		Notification: jrpc2.NewNotification(lsp.MethodShowMessage),
		Params: lsp.ShowMessageParams{
			Type:    lsp.Debug,
			Message: msg,
		},
	}
	send(w, not)
}
