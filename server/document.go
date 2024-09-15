package server

import (
	"log/slog"

	"github.com/scatternoodle/wflang/internal/lsp"
	"github.com/scatternoodle/wflang/wflang/lexer"
	"github.com/scatternoodle/wflang/wflang/parser"
)

func (srv *Server) updateDocument(doc lsp.TextDocumentItem) {
	srv.uri = doc.URI
	srv.parser = parser.New(lexer.New(doc.Text))
	var err error
	if srv.ast, err = srv.parser.AST(); err != nil {
		slog.Error("error retrieving new AST", "error", err, "parser errors", srv.parser.Errors())
	}
	srv.tokenEncoder = newTokenEncoder(srv.parser.Tokens())
	slog.Info("Document AST generated",
		"version", doc.Version,
		"uri", doc.URI,
		"number of tokens", len(srv.parser.Tokens()),
		"errors", len(srv.parser.Errors()),
	)
	srv.createSymbols()
}
