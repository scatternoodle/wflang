package lsp

// Text Document Completion
// https://microsoft.github.io/language-server-protocol/specifications/lsp/3.17/specification/#textDocument_completion

// ComletionOptions are the config options for the CompletionProvider sent in
// server capabilities.
//
// https://microsoft.github.io/language-server-protocol/specifications/lsp/3.17/specification/#completionOptions
type CompletionOptions struct {
	TriggerCharacters   []string               `json:"triggerCharacters,omitempty"`
	AllCommitCharacters []string               `json:"allCommitCharacters,omitempty"`
	ResolveProvider     bool                   `json:"resolveProvider,omitempty"`
	CompletionItem      *CompletionItemOptions `json:"completionItem,omitempty"`
}

type CompletionItemOptions struct {
	LabelDetailsSupport bool `json:"labelDetailsSupport,omitempty"`
}
