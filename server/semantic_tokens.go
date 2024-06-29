package server

import "github.com/scatternoodle/wflang/lsp"

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
