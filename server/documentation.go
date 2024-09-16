package server

import (
	"strings"

	"github.com/scatternoodle/wflang/internal/lsp"
	"github.com/scatternoodle/wflang/wflang/object"
	"github.com/scatternoodle/wflang/wflang/token"
)

func (srv *Server) hover(pos lsp.Position) lsp.Hover {
	_, tok, ok := srv.getTokenAtPos(pos)
	if !ok {
		return lsp.Hover{}
	}
	if tok.Type == token.T_BUILTIN {
		return lsp.Hover{MarkupContent: *object.DocMarkdown(strings.ToLower(tok.Literal))}
	}
	return lsp.Hover{}
}
