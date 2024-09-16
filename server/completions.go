package server

import (
	"github.com/scatternoodle/wflang/internal/lsp"
	"github.com/scatternoodle/wflang/wflang/object"
)

func (srv *Server) completions(pos lsp.Position) []lsp.CompletionItem {
	funcs := object.Builtins()
	items := make([]lsp.CompletionItem, 0, len(funcs))
	for _, fn := range funcs {
		label := fn.Name
		item := lsp.CompletionItem{
			Label: label,

			// LabelDetail omitted in favour of Documentation (less busy, doesn't get truncated)

			// Detail omitted, can be used in conjunction with Documentation but makes it all a bit busy
			// and we don't really have a need for this. Detail gets displayed above the doc string, but
			// in a completely different font and it's kinda weird (at least in VS****).

			Documentation:  object.DocMarkdown(label),
			Kind:           lsp.CompItemFunc,
			InsertText:     label,
			InsertFormat:   lsp.InsFormatPlainText,
			InsertTextMode: lsp.InsModeAsIs,
		}
		items = append(items, item)
	}
	return items
}
