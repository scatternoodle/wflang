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

func semanticTokenTypes() []string {
	return []string{
		// "namespace",
		semTokenType,
		semTokenClass,
		// "enum",
		// "interface",
		// "struct",
		// "typeParameter",
		semTokenParam,
		semTokenVar,
		semTokenProperty,
		semTokenEnumMember,
		// "event",
		semTokenFunc,
		semTokenMethod,
		semTokenMacro,
		semTokenKeyword,
		// "modifier",
		semTokenComment,
		semTokenString,
		semTokenNum,
		// "regexp",
		semTokenOperator,
		// "decorator",
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
