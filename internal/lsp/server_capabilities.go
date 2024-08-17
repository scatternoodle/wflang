package lsp

type ServerCapabilities struct {
	TextDocumentSync       TextDocumentSyncKind  `json:"textDocumentSync,omitempty"`
	SemanticTokensProvider SemanticTokensOptions `json:"semanticTokensProvider,omitempty"`
	HoverProvider          bool                  `json:"hoverProvider,omitempty"`
	DocumentSymbolProvider bool                  `json:"documentSymbolProvider,omitempty"`
	DeclarationProvider    bool                  `json:"declarationProvider,omitempty"`
	DefinitionProvider     bool                  `json:"definitionProvider,omitempty"`
}

type TextDocumentSyncKind int

const (
	SyncNone TextDocumentSyncKind = iota
	SyncFull
	SyncIncremental
)
