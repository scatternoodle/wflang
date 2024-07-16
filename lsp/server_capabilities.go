package lsp

type ServerCapabilities struct {
	TextDocumentSync       TextDocumentSyncKind  `json:"textDocumentSync,omitempty"`
	SemanticTokensProvider SemanticTokensOptions `json:"semanticTokensProvider,omitempty"`
	HoverProvider          bool                  `json:"hoverProvider,omitempty"`
}

type TextDocumentSyncKind int

const (
	SyncNone TextDocumentSyncKind = iota
	SyncFull
	SyncIncremental
)
