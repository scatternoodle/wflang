package server

import (
	"log/slog"
	"strings"

	"github.com/scatternoodle/wflang/lang/builtins"
	"github.com/scatternoodle/wflang/lang/token"
	"github.com/scatternoodle/wflang/lsp"
	"github.com/scatternoodle/wflang/server/hovertext"
)

func (srv *Server) hover(pos lsp.Position) lsp.Hover {
	hov := lsp.Hover{
		MarkupContent: lsp.MarkupContent{
			Kind:  lsp.MarkupKindMarkdown,
			Value: "",
		},
	}

	tok, ok := srv.getTokenAtPos(pos)
	if !ok {
		slog.Debug("No token found")
		return hov
	}
	slog.Debug("token found", "token", tok)

	if tok.Type == token.T_BUILTIN {
		hov.Value, _ = builtinHoverText(strings.ToLower(tok.Literal))
	}

	return hov
}

func builtinHoverText(name string) (text string, ok bool) {
	switch name {
	case builtins.SumTime:
		return hovertext.SumTime, true
	}

	return "", false
}
