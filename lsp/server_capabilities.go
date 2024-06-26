package lsp

type ServerCapabilities struct {
	TextDocumentSync       TextDocumentSyncKind  `json:"textDocumentSync,omitempty"`
	SemanticTokensProvider SemanticTokensOptions `json:"semanticTokensProvider,omitempty"`
}

type TextDocumentSyncKind int

const (
	SyncNone TextDocumentSyncKind = iota
	SyncFull
	SyncIncremental
)

type SemanticTokensOptions struct {
	Legend TokenTypesLegend `json:"legend"`
	Range  bool             `json:"range,omitempty"`
	Full   bool             `json:"full,omitempty"`
}

type TokenTypesLegend struct {
	TokenTypes     []string `json:"tokenTypes"`
	TokenModifiers []string `json:"tokenModifiers"`
}
