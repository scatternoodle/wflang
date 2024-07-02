package server

import (
	"fmt"
	"log/slog"

	"github.com/scatternoodle/wflang/lang/token"
	"github.com/scatternoodle/wflang/lsp"
)

func (srv *Server) getSemanticTokens() []lsp.Uint {
	slog.Info("Encoding semantic tokens...")
	pTokens := srv.parser.Tokens()
	sTokens := make([]lsp.Uint, 0, len(pTokens)*5)

	for i := range len(pTokens) {
		slog.Debug("encoding token", "i", i, "token", pTokens[i])
		encoded, ok := encodeToken(pTokens[i])
		if !ok {
			slog.Debug("unable to encode", "token", pTokens[i])
			continue
		}
		sTokens = append(sTokens, encoded...)
		slog.Debug("sTokens is now", "sTokens", sTokens)
	}

	slog.Info(fmt.Sprintf("Successfully encoded %d semantic tokens", len(sTokens)/5))
	return sTokens
}

// encodeToken encodes a Semantic Token according to the LSP specification. The
// token is represented by 5 integers:
//
//	{ line, startChar, length, tokenType, tokenModifiers }
//
// Token modifier indices are repesented by bit flags in the final integer, e.g. value
// 3 means tokenModifiers[0] and tokenModifiers[1]
func encodeToken(tk token.Token) (token []lsp.Uint, found bool) {
	out := make([]lsp.Uint, 5)
	tkType, ok := typeFromToken(tk.Type)
	if !ok {
		return out, false
	}
	out[3] = lsp.Uint(tkType)

	out[0] = lsp.Uint(tk.StartPos.Line)
	out[1] = lsp.Uint(tk.StartPos.Col)
	out[2] = lsp.Uint(tk.Len)

	// we'll implement modifier encoding if/when needed - we aren't using any
	// modifiers for now.
	out[4] = 0

	slog.Debug("encoded", "out", out)
	return out, true
}

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

type semanticTokenType lsp.Uint

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

// typeFromToken gets the server token type from the lexer token type given, or 0
// and false if not found.
func typeFromToken(t token.Type) (tType semanticTokenType, found bool) {
	switch t {
	case token.T_EOF:
		return semIndexVar, true
	case token.T_IDENT:
		return semIndexVar, true
	case token.T_NUM:
		return semIndexNum, true
	case token.T_STRING:
		return semIndexString, true
	case token.T_COMMENT_LINE:
		return semIndexComment, true
	case token.T_COMMENT_BLOCK:
		return semIndexComment, true
	case token.T_EQ:
		return semIndexOperator, true
	case token.T_PLUS:
		return semIndexOperator, true
	case token.T_MINUS:
		return semIndexOperator, true
	case token.T_BANG:
		return semIndexOperator, true
	case token.T_NEQ:
		return semIndexOperator, true
	case token.T_ASTERISK:
		return semIndexOperator, true
	case token.T_SLASH:
		return semIndexOperator, true
	case token.T_MODULO:
		return semIndexOperator, true
	case token.T_LT:
		return semIndexOperator, true
	case token.T_GT:
		return semIndexOperator, true
	case token.T_LTE:
		return semIndexOperator, true
	case token.T_GTE:
		return semIndexOperator, true
	case token.T_AND:
		return semIndexOperator, true
	case token.T_OR:
		return semIndexOperator, true
	case token.T_COMMA:
		return 0, false
	case token.T_SEMICOLON:
		return 0, false
	case token.T_COLON:
		return 0, false
	case token.T_LPAREN:
		return 0, false
	case token.T_RPAREN:
		return 0, false
	case token.T_LBRACE:
		return 0, false
	case token.T_RBRACE:
		return 0, false
	case token.T_LBRACKET:
		return 0, false
	case token.T_RBRACKET:
		return 0, false
	case token.T_PERIOD:
		return 0, false
	case token.T_DOLLAR:
		return 0, false
	case token.T_DOUBLEQUOTE:
		return 0, false
	case token.T_BUILTIN:
		return semIndexFunc, true
	case token.T_VAR:
		return semIndexKeyword, true
	case token.T_OVER:
		return semIndexKeyword, true
	case token.T_WHERE:
		return semIndexKeyword, true
	case token.T_ORDER:
		return semIndexKeyword, true
	case token.T_BY:
		return semIndexKeyword, true
	case token.T_ASC:
		return semIndexKeyword, true
	case token.T_DESC:
		return semIndexKeyword, true
	case token.T_ALIAS:
		return semIndexKeyword, true
	case token.T_IN:
		return semIndexKeyword, true
	case token.T_SET:
		return semIndexKeyword, true
	case token.T_NULL:
		return semIndexKeyword, true
	case token.T_TRUE:
		return semIndexKeyword, true
	case token.T_FALSE:
		return semIndexKeyword, true

	default:
		return 0, false
	}
}
