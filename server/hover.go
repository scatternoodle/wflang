package server

import (
	"fmt"
	"log/slog"

	"github.com/scatternoodle/wflang/lang/token"
	"github.com/scatternoodle/wflang/lsp"
)

func (srv *Server) hover(pos lsp.Position) lsp.Hover {
	hov := lsp.Hover{
		MarkupContent: lsp.MarkupContent{
			Kind:  lsp.MarkupKindMarkdown,
			Value: "",
		},
	}

	slog.Debug(fmt.Sprintf("Searching for token at %+v", pos))
	tok, ok := srv.getTokenAtPos(pos)
	if !ok {
		slog.Debug("No token found")
		return hov
	}

	slog.Debug(fmt.Sprintf("Token found: %+v", tok))
	if tok.Type == token.T_BUILTIN {
		hov.Value = "builtin"
	}

	return hov
}
