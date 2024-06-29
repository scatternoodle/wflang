package server

import (
	"github.com/scatternoodle/wflang/lang/token"
	"github.com/scatternoodle/wflang/lsp"
)

func semanticTokensProvider() lsp.SemanticTokensOptions {
	opt := lsp.SemanticTokensOptions{
		Legend: lsp.TokenTypesLegend{
			TokenTypes:     semanticTokenTypes(),
			TokenModifiers: semanticTokenModifiers(),
		},
		Range: false,
		Full:  true,
	}

	return opt
}

const (
	semTokenType       string = "type"
	semTokenClass      string = "class"
	semTokenParam      string = "parameter"
	semTokenVar        string = "variable"
	semTokenProperty   string = "property"
	semTokenEnumMember string = "enumMember"
	semTokenFunc       string = "function"
	semTokenMethod     string = "method"
	semTokenMacro      string = "macro"
	semTokenKeyword    string = "keyword"
	semTokenComment    string = "comment"
	semTokenString     string = "string"
	semTokenNum        string = "number"
	semTokenOperator   string = "operator"
)

const (
	semIndexType semanticTokenType = iota
	semIndexClass
	semIndexParam
	semIndexVar
	semIndexProperty
	semIndexEnumMember
	semIndexFunc
	semIndexMethod
	semIndexMacro
	semIndexKeyword
	semIndexComment
	semIndexString
	semIndexNum
	semIndexOperator
)

type semanticTokenType int32

func (s semanticTokenType) String() string {
	switch s {
	case semIndexType:
		return semTokenType
	case semIndexClass:
		return semTokenClass
	case semIndexParam:
		return semTokenParam
	case semIndexVar:
		return semTokenVar
	case semIndexProperty:
		return semTokenProperty
	case semIndexEnumMember:
		return semTokenEnumMember
	case semIndexFunc:
		return semTokenFunc
	case semIndexMethod:
		return semTokenMethod
	case semIndexMacro:
		return semTokenMacro
	case semIndexKeyword:
		return semTokenKeyword
	case semIndexComment:
		return semTokenComment
	case semIndexString:
		return semTokenString
	case semIndexNum:
		return semTokenNum
	case semIndexOperator:
		return semTokenOperator
	default:
		return ""
	}
}

func semanticTokenTypes() []string {
	return []string{
		semTokenType,
		semTokenClass,
		semTokenParam,
		semTokenVar,
		semTokenProperty,
		semTokenEnumMember,
		semTokenFunc,
		semTokenMethod,
		semTokenMacro,
		semTokenKeyword,
		semTokenComment,
		semTokenString,
		semTokenNum,
		semTokenOperator,
		// "namespace",
		// "enum",
		// "interface",
		// "struct",
		// "typeParameter",
		// "decorator",
		// "event",
		// "modifier",
		// "regexp",
	}
}

func semanticTokenModifiers() []string {
	return []string{ // We aren't yet supporting any, but these are the ones VSCode does...
		// "declaration",
		// "definition",
		// "readonly",
		// "static",
		// "deprecated",
		// "abstract",
		// "async",
		// "modification",
		// "documentation",
		// "defaultLibrary",
	}
}

// typeFromToken gets the server token type from the lexer token type given. If
// not mapped to an LSP Semantic Token, returns -1 and does not error (some lexer
// tokens don't have a Semantic Token in the LSP, and don't need one).
func typeFromToken(t token.Type) semanticTokenType {
	switch t {
	case token.T_EOF:
		return -1
	case token.T_IDENT:
		return semIndexVar
	case token.T_NUM:
		return semIndexNum
	case token.T_STRING:
		return semIndexString
	case token.T_COMMENT_LINE:
		return semIndexComment
	case token.T_COMMENT_BLOCK:
		return semIndexComment
	case token.T_EQ:
		return semIndexOperator
	case token.T_PLUS:
		return semIndexOperator
	case token.T_MINUS:
		return semIndexOperator
	case token.T_BANG:
		return semIndexOperator
	case token.T_NEQ:
		return semIndexOperator
	case token.T_ASTERISK:
		return semIndexOperator
	case token.T_SLASH:
		return semIndexOperator
	case token.T_MODULO:
		return semIndexOperator
	case token.T_LT:
		return semIndexOperator
	case token.T_GT:
		return semIndexOperator
	case token.T_LTE:
		return semIndexOperator
	case token.T_GTE:
		return semIndexOperator
	case token.T_AND:
		return semIndexOperator
	case token.T_OR:
		return semIndexOperator
	case token.T_COMMA:
		return -1
	case token.T_SEMICOLON:
		return -1
	case token.T_COLON:
		return -1
	case token.T_LPAREN:
		return -1
	case token.T_RPAREN:
		return -1
	case token.T_LBRACE:
		return -1
	case token.T_RBRACE:
		return -1
	case token.T_LBRACKET:
		return -1
	case token.T_RBRACKET:
		return -1
	case token.T_PERIOD:
		return -1
	case token.T_DOLLAR:
		return -1
	case token.T_DOUBLEQUOTE:
		return -1
	case token.T_BUILTIN:
		return semIndexFunc
	case token.T_VAR:
		return semIndexKeyword
	case token.T_OVER:
		return semIndexKeyword
	case token.T_WHERE:
		return semIndexKeyword
	case token.T_ORDER:
		return semIndexKeyword
	case token.T_BY:
		return semIndexKeyword
	case token.T_ASC:
		return semIndexKeyword
	case token.T_DESC:
		return semIndexKeyword
	case token.T_ALIAS:
		return semIndexKeyword
	case token.T_IN:
		return semIndexKeyword
	case token.T_SET:
		return semIndexKeyword
	case token.T_NULL:
		return semIndexKeyword
	case token.T_TRUE:
		return semIndexKeyword
	case token.T_FALSE:
		return semIndexKeyword

	default:
		return -1
	}
}
