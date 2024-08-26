package lsp

// Signature Help
//
// https://microsoft.github.io/language-server-protocol/specifications/lsp/3.17/specification/#textDocument_signatureHelp

type SignatureHelpOptions struct {
	TriggerChars   []string `json:"triggerCharacters,omitempty"`
	RetriggerChars []string `json:"retriggerCharacters,omitempty"`
}
