package server

import (
	"io"
	"log/slog"
	"os"

	"github.com/scatternoodle/wflang/jrpc2"
	"github.com/scatternoodle/wflang/lsp"
)

// handlerFunc takes an io.Writer and a byte slice containing the contents of an
// LSP Request or Notification, processes, and responds accordingly. For
// Notifications, id can be nil.
type handlerFunc func(w io.Writer, c []byte, id *int)

func (srv *Server) handleInitializeRequest(w io.Writer, c []byte, id *int) {
	var r lsp.InitializeRequest
	if !handleAssertID(w, id) || !handleParseContent(&r, w, c, id) {
		return
	}
	respond(w, srv.initialize(id))
}

func (srv *Server) handleInitializedNotification(w io.Writer, c []byte, id *int) {
	srv.initialized = true
}

func (srv *Server) handleShutdownRequest(w io.Writer, c []byte, id *int) {
	if !handleAssertID(w, id) {
		return
	}
	srv.exiting = true
	respond(w, struct {
		jrpc2.Response
		Result any `json:"result"`
	}{
		Response: jrpc2.NewResponse(id, nil),
		Result:   nil,
	})
}

func (srv *Server) handleExitNotification(_ io.Writer, _ []byte, _ *int) {
	var errCode int
	if !srv.exiting {
		errCode = 1
	}
	slog.Info("Server exiting", "code", errCode)
	os.Exit(errCode)
}

func (srv *Server) handleDocDidOpenNotification(w io.Writer, c []byte, id *int) {
	var r lsp.NotificationDidOpen
	if !handleParseContent(&r, w, c, id) {
		return
	}
	srv.updateDocument(r.Params.TextDocument)
}

func (srv *Server) handleSemanticTokensFullRequest(w io.Writer, c []byte, id *int) {
	var r lsp.SemanticTokensRequest
	if !handleAssertID(w, id) || !handleParseContent(&r, w, c, id) {
		return
	}
	respond(w, &lsp.SemanticTokensResponse{
		Response: jrpc2.NewResponse(id, nil),
		Result: lsp.SemanticTokensResult{
			Data: srv.semTokens,
		},
	})
}
