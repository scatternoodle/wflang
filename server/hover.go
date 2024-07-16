package server

import (
	"github.com/scatternoodle/wflang/lang/token"
	"github.com/scatternoodle/wflang/lsp"
)

func (srv *Server) hover(pos lsp.Position) lsp.Hover {
	hov := lsp.Hover{
		MarkupContent: lsp.MarkupContent{
			Kind:  lsp.MarkupKindMarkdown,
			Value: "",
		},
		Position: pos,
	}

	tok, ok := srv.getTokenAtPos(pos)
	if !ok {
		return hov
	}

	if tok.Type == token.T_BUILTIN {
		hov.Value = "builtin"
	}

	return hov
}
