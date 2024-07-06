package server

import (
	"log/slog"
	"slices"

	"github.com/scatternoodle/wflang/lang/token"
)

// See Microsoft LSP spec for detailed explanation on semantic token encoding.
// https://microsoft.github.io/language-server-protocol/specifications/lsp/3.17/specification/#textDocument_semanticTokens

// All LSP semantic token types:
//  type = 'type',
// 	class = 'class',
// 	enum = 'enum',
// 	interface = 'interface',
// 	struct = 'struct',
// 	typeParameter = 'typeParameter',
// 	parameter = 'parameter',
// 	variable = 'variable',
// 	property = 'property',
// 	enumMember = 'enumMember',
// 	event = 'event',
// 	function = 'function',
// 	method = 'method',
// 	macro = 'macro',
// 	keyword = 'keyword',
// 	modifier = 'modifier',
// 	comment = 'comment',
// 	string = 'string',
// 	number = 'number',
// 	regexp = 'regexp',
// 	operator = 'operator',
// 	/**
// 	 * @since 3.17.0
// 	 */
// 	decorator = 'decorator'

const (
	semKeyword = "keyword"
)

func newTokenEncoder(tokens []token.Token) *tokenEncoder {
	slog.Debug("newTokenEncoder called with", "tokens", tokens)
	e := &tokenEncoder{
		types: []string{
			semKeyword,
		},
		typeMap: map[token.Type]string{
			token.T_BUILTIN: semKeyword,
			token.T_VAR:     semKeyword,
		},
		semTokens: []uint{},
	}
	e.encode(tokens)
	return e
}

// tokenEncoder stores encoded LSP semantic tokens, as well as the type legend for
// the language server.
type tokenEncoder struct {
	types     []string
	semTokens []uint
	typeMap   map[token.Type]string
	// modifiers []string - not currently handled.
}

func (t *tokenEncoder) encode(parserTokens []token.Token) {
	var (
		prvLine, prvCol     uint
		crrLine, crrCol     uint
		deltaLine, deltaCol uint
	)
	semTok := make([]uint, 5)

	for _, token := range parserTokens {
		typeStr, ok := t.typeMap[token.Type]
		if !ok {
			continue
		}

		idx := slices.Index(t.types, typeStr)
		if idx < 0 {
			slog.Error(
				"Type string in tokenEncoder typeMap but not in registered types array.",
				"typeStr", typeStr,
				"token.Type", token.Type)
			continue
		}

		crrLine = uint(token.StartPos.Line)
		crrCol = uint(token.StartPos.Col)
		deltaLine = crrLine - prvLine
		deltaCol = crrCol
		if crrLine == prvLine {
			deltaCol -= prvCol
		}
		prvLine = crrLine
		prvCol = crrCol

		semTok[0] = deltaLine
		semTok[1] = deltaCol
		semTok[2] = uint(token.Len)
		semTok[3] = uint(idx)
		semTok[4] = 0 // not currently handling modifier bitmasks

		t.semTokens = append(t.semTokens, semTok...)
	}
}
