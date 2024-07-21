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
	case builtins.Min:
		return hovertext.Min, true
	case builtins.Max:
		return hovertext.Max, true
	case builtins.Contains:
		return hovertext.Contains, true
	case builtins.Sum:
		return hovertext.Sum, true
	case builtins.Count:
		return hovertext.Count, true
	case builtins.If:
		return hovertext.If, true
	case builtins.SumTime:
		return hovertext.SumTime, true
	case builtins.CountTime:
		return hovertext.CountTime, true
	case builtins.FindFirstTime:
		return hovertext.FindFirstTime, true
	case builtins.SumSchedule:
		return hovertext.SumSchedule, true
	case builtins.CountSchedule:
		return hovertext.CountSchedule, true
	case builtins.FindFirstSchedule:
		return hovertext.FindFirstSchedule, true
	case builtins.CountException:
		return hovertext.CountException, true
	}

	return "", false
}
