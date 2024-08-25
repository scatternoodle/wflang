package lsp

// Command - careful - commands are client-side and not known to the protocol.
//
// https://microsoft.github.io/language-server-protocol/specifications/lsp/3.17/specification/#command
type Command struct {
	Title   string `json:"title"`
	Command string `json:"command"`
	Args    []any  `json:"arguments,omitempty"`
}
