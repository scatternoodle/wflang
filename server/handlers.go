package server

import (
	"io"
	"log/slog"
	"os"

	"github.com/scatternoodle/wflang/internal/jrpc2"
	"github.com/scatternoodle/wflang/internal/lsp"
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

func (srv *Server) handleInitializedNotification(_ io.Writer, _ []byte, _ *int) {
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

func (srv *Server) handleDocDidChangeNotification(w io.Writer, c []byte, id *int) {
	var r lsp.NotificationDidChange
	if !handleParseContent(&r, w, c, id) {
		return
	}

	// we will update document with the last item in the changes array - we're only
	// only interested in the latest state of the document as we do not support
	// incremental updates.
	changes := r.Params.ContentChanges
	lastChange := changes[len(changes)-1]

	srv.updateDocument(
		lsp.TextDocumentItem{
			URI:     r.Params.TextDocument.URI,
			Version: r.Params.TextDocument.Version,
			Text:    lastChange.Text,
		})
}

func (srv *Server) handleDocDidSaveNotification(w io.Writer, c []byte, id *int) {
	// currently no-op
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

func (srv *Server) handleHoverRequest(w io.Writer, c []byte, id *int) {
	var r lsp.HoverRequest
	if !handleAssertID(w, id) || !handleParseContent(&r, w, c, id) {
		return
	}

	respond(w, lsp.HoverResponse{
		Response: jrpc2.NewResponse(id, nil),
		Hover:    srv.hover(r.Position),
	})
}

func (srv *Server) handleDocumentSymbolsRequest(w io.Writer, c []byte, id *int) {
	var r lsp.DocumentSymbolRequest
	if !handleAssertID(w, id) || !handleParseContent(&r, w, c, id) {
		return
	}

	res := map[string]lsp.DocumentSymbol{}
	if srv.symbols != nil {
		res = srv.symbols
	}
	respond(w, lsp.DocumentSymbolResponse{
		Response: jrpc2.NewResponse(id, nil),
		Result:   res,
	})
}

func (srv *Server) handleGotoDeclarationRequest(w io.Writer, c []byte, id *int) {
	var r lsp.GotoDeclarationRequest
	if !handleAssertID(w, id) || !handleParseContent(&r, w, c, id) {
		return
	}
	respond(w, lsp.GotoDeclarationRequest{}) // TODO
}

func (srv *Server) handleGotoDefinitionRequest(w io.Writer, c []byte, id *int) {
	var r lsp.GotoDefinitionRequest
	if !handleAssertID(w, id) || !handleParseContent(&r, w, c, id) {
		return
	}
	_ = r
}
