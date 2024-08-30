package lsp

import "github.com/scatternoodle/wflang/internal/jrpc2"

// SymbolKind
//
// https://microsoft.github.io/language-server-protocol/specifications/lsp/3.17/specification/#symbolKind
type SymbolKind int

const (
	_ SymbolKind = iota
	SYMBOL_KIND_FILE
	SYMBOL_KIND_MODULE
	SYMBOL_KIND_NAMESPACE
	SYMBOL_KIND_PACKAGE
	SYMBOL_KIND_CLASS
	SYMBOL_KIND_METHOD
	SYMBOL_KIND_PROPERTY
	SYMBOL_KIND_FIELD
	SYMBOL_KIND_CONSTRUCTOR
	SYMBOL_KIND_ENUM
	SYMBOL_KIND_INTERFACE
	SYMBOL_KIND_FUNCTION
	SYMBOL_KIND_VARIABLE
	SYMBOL_KIND_CONSTANT
	SYMBOL_KIND_STRING
	SYMBOL_KIND_NUMBER
	SYMBOL_KIND_BOOLEAN
	SYMBOL_KIND_ARRAY
	SYMBOL_KIND_OBJECT
	SYMBOL_KIND_KEY
	SYMBOL_KIND_NULL
	SYMBOL_KIND_ENUM_MEMBER
	SYMBOL_KIND_STRUCT
	SYMBOL_KIND_EVENT
	SYMBOL_KIND_OPERATOR
	SYMBOL_KIND_TYPE_PARAMETER
)

// SymbolTag is an extra annotation that editors can use to render a symbol in a
// special way.
//
// https://microsoft.github.io/language-server-protocol/specifications/lsp/3.17/specification/#symbolTag
type SymbolTag int

const SYMTAG_DEPRECATED SymbolTag = 1

// DocumentSymbol represents programming constructs like variables, classes, etc.
// that appear in a document.
//
// https://microsoft.github.io/language-server-protocol/specifications/lsp/3.17/specification/#documentSymbol
type DocumentSymbol struct {
	Name string `json:"name"`

	// Detail optional info to render with the symbol, e.g. a function signature.
	Detail string      `json:"detail,omitempty"`
	Kind   SymbolKind  `json:"kind"`
	Tags   []SymbolTag `json:"tags,omitempty"`

	// Deprecated: use the SymbolTag SYMTAG_DEPRECATED in Tags instead.
	Deprecated bool `json:"deprecated,omitempty"`

	// Range encompassing where the range you want the cursor to be able to pick
	// up the symbol. Can be larger than the SelectionRange, which strictly
	// encompasses the symbol itself.
	Range Range `json:"range"`

	// SelectionRange encompasses the symbol itself, and is used to highlight the
	// symbol in the editor / drive selection.
	SelectionRange Range `json:"selectionRange"`

	// Children is a list of child symbols that are part of this symbol, e.g.
	// methods in a class.
	Children []DocumentSymbol `json:"children,omitempty"`
}

// DocumentSymbolRequest is sent from the client to the server to request a list
// of all symbols in a document.
//
// https://microsoft.github.io/language-server-protocol/specifications/lsp/3.17/specification/#textDocument_documentSymbol
type DocumentSymbolRequest struct {
	jrpc2.Request
	Params DocumentSymbolParams `json:"params"`
}

// DocumentSymbolParams
//
// https://microsoft.github.io/language-server-protocol/specifications/lsp/3.17/specification/#documentSymbolParams
type DocumentSymbolParams struct {
	TextDocument TextDocumentIdentifier `json:"textDocument"`
}

// DocumentSymbolResponse is sent from the server to the client in response to a
// DocumentSymbolRequest.
//
// https://microsoft.github.io/language-server-protocol/specifications/lsp/3.17/specification/#textDocument_documentSymbol
type DocumentSymbolResponse struct {
	jrpc2.Response
	Result []DocumentSymbol `json:"result"`
}

// GotoDefinitionRequest is sent from the client to the server to resolve the
// definition location of a symbol at a given text document position.
//
// https://microsoft.github.io/language-server-protocol/specifications/lsp/3.17/specification/#textDocument_definition
type GotoDefinitionRequest struct {
	jrpc2.Request
	Params TextDocumentPositionParams `json:"params"`
}

// GotoDefinitionResponse
//
// https://microsoft.github.io/language-server-protocol/specifications/lsp/3.17/specification/#textDocument_definition
type GotoDefinitionResponse struct {
	jrpc2.Response
	Result *Location `json:"result"`
}
