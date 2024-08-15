package server

import (
	"github.com/scatternoodle/wflang/lsp"
)

func (srv *Server) documentSymbols() []lsp.DocumentSymbol {
	symbols := []lsp.DocumentSymbol{}

	for _, v := range srv.parser.Vars() {
		if v.Statement == nil {
			continue
		}

		varStart, varEnd := v.Statement.Pos()
		symbolRange := lsp.Range{
			Start: lsp.Position{Line: varStart.Line, Character: varStart.Col},
			End:   lsp.Position{Line: varEnd.Line, Character: varEnd.Col},
		}

		nameStart, nameEnd := v.Statement.Name.Pos()
		selectionRange := lsp.Range{
			Start: lsp.Position{Line: nameStart.Line, Character: nameStart.Col},
			End:   lsp.Position{Line: nameEnd.Line, Character: nameEnd.Col},
		}

		symbols = append(symbols, lsp.DocumentSymbol{
			Name:           v.Name,
			Kind:           lsp.SYMBOL_KIND_VARIABLE,
			Range:          symbolRange,
			SelectionRange: selectionRange,
		})
	}
	return symbols
}
