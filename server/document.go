package server

import (
	"log/slog"

	"github.com/scatternoodle/wflang/internal/lsp"
	"github.com/scatternoodle/wflang/lang/lexer"
	"github.com/scatternoodle/wflang/lang/parser"
)

func (srv *Server) updateDocument(doc lsp.TextDocumentItem) {
	srv.parser = parser.New(lexer.New(doc.Text))
	srv.tokenEncoder = newTokenEncoder(srv.parser.Tokens())

	slog.Info("Document AST generated",
		"version", doc.Version,
		"uri", doc.URI,
		"number of tokens", len(srv.parser.Tokens()),
		"errors", len(srv.parser.Errors()),
	)

	srv.createSymbols()
}
