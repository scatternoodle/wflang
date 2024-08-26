package lsp

import "github.com/scatternoodle/wflang/internal/jrpc2"

// Rename
//
// https://microsoft.github.io/language-server-protocol/specifications/lsp/3.17/specification/#textDocument_rename

// RenameRequest
//
// https://microsoft.github.io/language-server-protocol/specifications/lsp/3.17/specification/#textDocument_rename
type RenameRequest struct {
	jrpc2.Request
	RenameParams `json:"params"`
}

// RenameParams
//
// https://microsoft.github.io/language-server-protocol/specifications/lsp/3.17/specification/#textDocument_rename
type RenameParams struct {
	TextDocumentPositionParams
	NewName string `json:"newName"`
}

// RenameResponse
//
// https://microsoft.github.io/language-server-protocol/specifications/lsp/3.17/specification/#textDocument_rename
type RenameResponse struct {
	jrpc2.Response
	Result *WorkspaceEdit `json:"result,omitempty"`
}
