package server

import (
	"github.com/scatternoodle/wflang/internal/lsp"
	"github.com/scatternoodle/wflang/lang/token"
)

func (srv *Server) createSymbols() {
	srv.symbols = map[string]lsp.DocumentSymbol{}

	for _, v := range srv.parser.Vars() {
		if v.Statement == nil {
			continue
		}

		varStart, varEnd := v.Statement.Pos()

		// For some reason (at least in VSCode) the end (and the end only) of the larger range in a document
		// symbol behaves like it's 1-indexed instead of zero-indexed
		varEnd.Col++

		symbolRange := lsp.Range{
			Start: lsp.Position{Line: varStart.Line, Character: varStart.Col},
			End:   lsp.Position{Line: varEnd.Line, Character: varEnd.Col},
		}

		nameStart, nameEnd := v.Statement.Name.Pos()
		selectionRange := lsp.Range{
			Start: lsp.Position{Line: nameStart.Line, Character: nameStart.Col},
			End:   lsp.Position{Line: nameEnd.Line, Character: nameEnd.Col + 1},
		}

		srv.symbols[v.Name] = lsp.DocumentSymbol{
			Name:           v.Name,
			Kind:           lsp.SYMBOL_KIND_VARIABLE,
			Range:          symbolRange,
			SelectionRange: selectionRange,
		}
	}
}

func (srv *Server) symbolFromPos(pos lsp.Position) (lsp.DocumentSymbol, bool) {
	_, tok, ok := srv.getTokenAtPos(pos)
	if !ok {
		return lsp.DocumentSymbol{}, false
	}
	if tok.Type != token.T_IDENT {
		return lsp.DocumentSymbol{}, false
	}

	sym, ok := srv.symbols[tok.Literal]
	if !ok {
		return lsp.DocumentSymbol{}, false
	}
	return sym, true
}
