package server

import (
	"fmt"
	"io"
	"log/slog"
	"os"
	"slices"

	"github.com/scatternoodle/wflang/internal/jrpc2"
	"github.com/scatternoodle/wflang/internal/lsp"
	"github.com/scatternoodle/wflang/lang/token"
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
	if err := srv.setInit(r.Params); err != nil {
		slog.Error("processing client initialize request failed", "error", err)
		respondError(w, id, lsp.ERRCODE_REQUEST_FAILED, err.Error())
	}
	send(w, srv.initializeResponse(id))
}

func (srv *Server) handleInitializedNotification(w io.Writer, _ []byte, _ *int) {
	srv.initialized = true
}

func (srv *Server) handleShutdownRequest(w io.Writer, c []byte, id *int) {
	if !handleAssertID(w, id) {
		return
	}
	srv.exiting = true
	send(w, struct {
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

func (srv *Server) handleSetTraceNotification(w io.Writer, c []byte, id *int) {
	var r lsp.SetTraceNotification
	if !handleParseContent(&r, w, c, id) {
		return
	}
	if err := srv.setTrace(r.TraceValue); err != nil {
		slog.Error("unable to set trace", "error", err)
		respondError(w, id, lsp.ERRCODE_REQUEST_FAILED, err.Error())
	}
	srv.trace = r.TraceValue
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
	send(w, &lsp.SemanticTokensResponse{
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

	send(w, lsp.HoverResponse{
		Response: jrpc2.NewResponse(id, nil),
		Hover:    srv.hover(r.Position),
	})
}

func (srv *Server) handleDocumentSymbolsRequest(w io.Writer, c []byte, id *int) {
	var r lsp.DocumentSymbolRequest
	if !handleAssertID(w, id) || !handleParseContent(&r, w, c, id) {
		return
	}

	res := []lsp.DocumentSymbol{}
	if srv.symbols != nil {
		for _, sym := range srv.symbols {
			res = append(res, sym)
		}
	}
	send(w, lsp.DocumentSymbolResponse{
		Response: jrpc2.NewResponse(id, nil),
		Result:   res,
	})
}

func (srv *Server) handleGotoDefinitionRequest(w io.Writer, c []byte, id *int) {
	var reqObj lsp.GotoDefinitionRequest
	if !handleAssertID(w, id) || !handleParseContent(&reqObj, w, c, id) {
		return
	}

	res := lsp.GotoDefinitionResponse{Response: jrpc2.NewResponse(id, nil)}
	sym, ok := srv.symbolFromPos(reqObj.Params.Position)
	if ok {
		res.Result = &lsp.Location{
			URI:   reqObj.Params.URI,
			Range: sym.SelectionRange,
		}
	}
	send(w, res)
}

func (srv *Server) handleCompletionRequest(w io.Writer, c []byte, id *int) {
	var req lsp.CompletionRequest
	if !handleAssertID(w, id) || !handleParseContent(&req, w, c, id) {
		return
	}
	send(w, lsp.CompletionResponse{
		Response: jrpc2.NewResponse(id, nil),
		Result:   srv.completions(req.Params.Position),
	})
}

func (srv *Server) handleRenameRequest(w io.Writer, c []byte, id *int) {
	var req lsp.RenameRequest
	if !handleAssertID(w, id) || !handleParseContent(&req, w, c, id) {
		return
	}

	_, reqTok, ok := srv.getTokenAtPos(req.Position)
	if !ok {
		respondError(w, id, lsp.ERRCODE_REQUEST_FAILED, fmt.Sprintf("no token found at position %+v", req.Position), nil)
	}
	if reqTok.Type != token.T_IDENT {
		respondError(w, id, lsp.ERRCODE_REQUEST_FAILED, fmt.Sprintf("invalid token type for rename %s", reqTok.Type), nil)
	}

	vars := srv.parser.Vars()
	varNames := make([]string, 0, len(vars))
	for _, varObj := range vars {
		varNames = append(varNames, varObj.Name)
	}
	if !slices.Contains(varNames, reqTok.Literal) {
		respondError(w, id, lsp.ERRCODE_REQUEST_FAILED, fmt.Sprintf("token %s is not a variable", reqTok.Literal), nil)
	}

	edits := []lsp.TextEdit{}
	for _, tok := range srv.parser.Tokens() {
		if tok.Literal != reqTok.Literal {
			continue
		}
		edits = append(edits, lsp.TextEdit{
			Range: lsp.Range{
				Start: lsp.Position{
					Line:      tok.StartPos.Line,
					Character: tok.StartPos.Col,
				},
				End: lsp.Position{
					Line:      tok.EndPos.Line,
					Character: tok.EndPos.Col + 1,
				},
			},
			NewText: req.NewName,
		})
	}
	wsEdit := &lsp.WorkspaceEdit{Changes: map[string][]lsp.TextEdit{srv.uri: edits}}
	send(w, lsp.RenameResponse{
		Response: jrpc2.NewResponse(id, nil),
		Result:   wsEdit,
	})
}

func (srv *Server) handleSignatureHelpRequest(w io.Writer, c []byte, id *int) {
	var req lsp.SignatureHelpRequest
	if !handleAssertID(w, id) || !handleParseContent(&req, w, c, id) {
		return
	}
	resp := lsp.SignatureHelpResponse{
		Response:      jrpc2.NewResponse(id, nil),
		SignatureHelp: nil,
	}

	idx, tkn, ok := srv.getTokenAtPos(cursorPos(req.Position))
	if !ok {
		send(w, resp)
	}
	_ = tkn
	if idx > 0 {
		idx--
		if idx >= len(srv.parser.Tokens()) {
			respondError(w, id, lsp.ERRCODE_REQUEST_FAILED,
				fmt.Sprintf("token idx %d is out of bounds with token array length %d", idx, len(srv.parser.Tokens())))
		}
		tkn = srv.parser.Tokens()[idx]
	}

	if tkn.Type != token.T_BUILTIN {
		send(w, resp)
	}

	send(w, resp)
}
