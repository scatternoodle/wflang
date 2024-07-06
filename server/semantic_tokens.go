package server

import (
	"log/slog"
	"slices"

	"github.com/scatternoodle/wflang/lang/token"
)

// See Microsoft LSP spec for detailed explanation on semantic token encoding.
// https://microsoft.github.io/language-server-protocol/specifications/lsp/3.17/specification/#textDocument_semanticTokens

const (
	// semType string = "type"
	// semClass string = "class"
	// semEnum string = "enum"
	// semInterface string = "interface"
	// semStruct string = "struct"
	// semTypeParameter string = "typeParameter"
	// semParameter string = "parameter"
	semVariable   string = "variable"
	semProperty   string = "property"
	semEnumMember string = "enumMember"
	// semEvent string = "event"
	semFunction string = "function"
	semMethod   string = "method"
	semMacro    string = "macro"
	semKeyword  string = "keyword"
	// semModifier string = "modifier"
	semComment string = "comment"
	semString  string = "string"
	semNumber  string = "number"
	// semRegexp string = "regexp"
	semOperator  string = "operator"
	semDecorator string = "decorator"
)

func tokenTypes() []string {
	return []string{
		// semType,
		// semClass,
		// semEnum,
		// semInterface,
		// semStruct,
		// semTypeParameter,
		// semParameter,
		semVariable,
		semProperty,
		semEnumMember,
		// semEvent,
		semFunction,
		semMethod,
		semMacro,
		semKeyword,
		// semModifier,
		semComment,
		semString,
		semNumber,
		// semRegexp,
		semOperator,
		semDecorator,
	}
}

func tokenModifiers() []string {
	return []string{}
}

func tokenMap() map[token.Type]string {
	return map[token.Type]string{
		token.T_COMMENT_BLOCK: semComment,
		token.T_COMMENT_LINE:  semComment,

		token.T_NUM:    semNumber,
		token.T_STRING: semString,

		token.T_BUILTIN: semFunction,

		token.T_VAR:   semKeyword,
		token.T_OVER:  semKeyword,
		token.T_WHERE: semKeyword,
		token.T_ORDER: semKeyword,
		token.T_BY:    semKeyword,
		token.T_ASC:   semKeyword,
		token.T_DESC:  semKeyword,
		token.T_ALIAS: semKeyword,
		token.T_IN:    semKeyword,
		token.T_SET:   semKeyword,
		token.T_NULL:  semKeyword,
		token.T_TRUE:  semKeyword,
		token.T_FALSE: semKeyword,
	}
}

func newTokenEncoder(tokens []token.Token) *tokenEncoder {
	slog.Debug("newTokenEncoder called with", "tokens", tokens)
	e := &tokenEncoder{
		types:   tokenTypes(),
		typeMap: tokenMap(),
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
