package server

import (
	"github.com/scatternoodle/wflang/internal/lsp"
	"github.com/scatternoodle/wflang/lang/object"
)

func (srv *Server) Completions(pos lsp.Position) []lsp.CompletionItem {
	funcs := object.Builtins()
	items := make([]lsp.CompletionItem, 0, len(funcs))
	for _, fn := range funcs {
		label := fn.Name
		item := lsp.CompletionItem{
			Label: label,
			// LabelDetails:
			Kind:           lsp.CompItemFunc,
			InsertText:     label,
			InsertFormat:   lsp.InsFormatPlainText,
			InsertTextMode: lsp.InsModeAsIs,
		}
		items = append(items, item)
	}
	return items
}
