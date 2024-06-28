package server

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"log/slog"

	"github.com/scatternoodle/wflang/jrpc2"
	"github.com/scatternoodle/wflang/lsp"
)

func New(name, version *string) *Server {
	return &Server{
		name:         name,
		version:      version,
		initialized:  false,
		capabilities: serverCapabilities(),
	}
}

type Server struct {
	name         *string
	version      *string
	initialized  bool
	capabilities lsp.ServerCapabilities
}

func serverCapabilities() lsp.ServerCapabilities {
	return lsp.ServerCapabilities{
		TextDocumentSync: lsp.SyncFull,
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
	slog.Debug(fmt.Sprintf("content: %s", string(content)))

	if !srv.initialized && method != lsp.MethodInitialize && method != lsp.MethodInitialized {
		respondError(w, requestId, lsp.ERRCODE_SERVER_NOT_INITIALIZED, "server not yet initialized", nil)
	}

	switch method {
	case lsp.MethodInitialize:
		var initReq lsp.InitializeRequest
		if err := json.Unmarshal(content, &initReq); err != nil {
			slog.Error("can't marshal request", "error", err)
			return
		}
		if initReq.ID == nil {
			slog.Error("request ID is nil")
			return
		}

		respond(w, srv.initialize(initReq.ID))
		slog.Info("InitializeResponse sent")

	case lsp.MethodInitialized:
		srv.initialized = true

	default:
		slog.Warn("method not handled", "method", method, "id", requestId)
	}

}

func respond(w io.Writer, v any) {
	response, err := jrpc2.EncodeMessage(v)
	if err != nil {
		panic(fmt.Errorf("response failed during encoding: %w", err))
	}
	if _, err = w.Write([]byte(response)); err != nil {
		panic(fmt.Errorf("response failed during write: %w", err))
	}
	slog.Debug("wrote response", "content", response)
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
