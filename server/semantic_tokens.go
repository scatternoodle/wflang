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
	for _, token := range parserTokens {
		typeStr, ok := t.typeMap[token.Type]
		if !ok {
			continue
		}

		idx := slices.Index(t.types, typeStr)
		if idx < 0 {
			slog.Error(
				"Type string in tokenEncoder typeMap but not in registered types array", "typeStr", typeStr, "token.Type", token.Type)
			continue
		}

		t.semTokens = append(t.semTokens, t.encodeToken(token, uint(idx))...)
	}
}

// encodeToken takes a token.Token and returns it as an token consists of 5 uints:
// array of uints encoded according to the LSP semanticToken structure:
//
//	{ line, startChar, length, tokenType, tokenModifiers }
func (t *tokenEncoder) encodeToken(tok token.Token, tType uint) []uint {
	return []uint{
		uint(tok.StartPos.Line),
		uint(tok.StartPos.Col),
		uint(tok.Len),
		tType,
		0,
	}
}
