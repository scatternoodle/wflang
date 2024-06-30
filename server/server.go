package server

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"log/slog"
	"os"

	"github.com/scatternoodle/wflang/jrpc2"
	"github.com/scatternoodle/wflang/lang/parser"
	"github.com/scatternoodle/wflang/lsp"
)

func New(name, version *string) *Server {
	return &Server{
		name:         name,
		version:      version,
		initialized:  false,
		capabilities: serverCapabilities(),
		parser:       nil,
	}
}

type Server struct {
	name         *string
	version      *string
	capabilities lsp.ServerCapabilities
	initialized  bool // before this is set true, we only accept requests with initialize method
	exiting      bool // set after an shutdown request is received, awaiting exit request
	parser       *parser.Parser
}

func serverCapabilities() lsp.ServerCapabilities {
	return lsp.ServerCapabilities{
		TextDocumentSync:       lsp.SyncFull,
		SemanticTokensProvider: semanticTokensProvider(),
	}
}

func (srv *Server) initialize(id *int32) lsp.InitializeResponse {
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
	slog.Info("Recieved", "method", method, "id", requestId)
	slog.Debug(fmt.Sprintf("Content=%s", string(content)))

	if !srv.initialized && method != lsp.MethodInitialize && method != lsp.MethodInitialized {
		respondError(w, requestId, lsp.ERRCODE_SERVER_NOT_INITIALIZED, "server not yet initialized", nil)
		return
	}

	if srv.exiting && method != lsp.MethodShutdown && method != lsp.MethodExit {
		respondError(w, requestId, jrpc2.ERRCODE_INVALID_REQUEST, "server is shutting down, expects 'exit' method", nil)
		return
	}

	switch method {
	case lsp.MethodInitialize:
		var initReq lsp.InitializeRequest
		if err := json.Unmarshal(content, &initReq); err != nil {
			slog.Error("Can't marshal request", "error", err)
			return
		}
		if initReq.ID == nil {
			slog.Error("Request ID is nil")
			return
		}

		respond(w, srv.initialize(initReq.ID))
		slog.Info("InitializeResponse sent")

	case lsp.MethodInitialized:
		srv.initialized = true

	case lsp.MethodDocDidOpen:
		var req lsp.NotificationDidOpen
		if err := json.Unmarshal(content, &req); err != nil {
			slog.Error("Can't marshal request", "error", err)
			return
		}
		srv.updateDocument(req.Params.TextDocument)

	case lsp.MethodSemanticTokensFull:
		var req lsp.RequestSemanticTokensFull
		if err := json.Unmarshal(content, &req); err != nil {
			slog.Error("Can't marshall request", "error", err)
			return
		}
		if req.ID == nil {
			slog.Error("Request ID is nil")
			return
		}

		slog.Debug("Parsed semantic tokens request", "object", req)

	case lsp.MethodShutdown:
		srv.exiting = true
		resp := struct {
			jrpc2.Response
			Result any `json:"result"`
		}{
			Response: jrpc2.NewResponse(requestId, nil),
			Result:   nil,
		}
		respond(w, &resp)

	case lsp.MethodExit:
		errCode := 0
		if !srv.exiting {
			errCode = 1
		}
		slog.Info("Server exiting", "code", errCode)
		os.Exit(errCode)

	default:
		slog.Warn("Unhandled method", "method", method, "id", requestId)
	}

}

func respond(w io.Writer, v any) {
	response, err := jrpc2.EncodeMessage(v)
	if err != nil {
		panic(fmt.Errorf("response failed during encoding: %w", err))
	}
	if _, err = w.Write(response); err != nil {
		panic(fmt.Errorf("response failed during write: %w", err))
	}
	slog.Debug(fmt.Sprintf("Wrote content=%s", string(response)))
}

func respondError(w io.Writer, id *int32, code int32, msg string, dat any) {
	rErr := jrpc2.ResponseError{
		Code:    code,
		Message: msg,
		Data:    dat,
	}
	v := jrpc2.NewResponse(id, &rErr)
	respond(w, v)
}

// getRequestID returns the Request ID from a content byte slice. Returns null if
// unable to resolve the Request ID.
func getRequestID(b []byte) *int32 {
	var idObj struct {
		ID *int32 `json:"id,omitempty"` // We want to be pretty permissive here as we can just return null if we can't find
	} // the ID.
	if err := json.Unmarshal(b, &idObj); err != nil {
		return nil
	}
	return idObj.ID
}
