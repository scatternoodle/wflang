package lsp

import "github.com/scatternoodle/wflang/internal/jrpc2"

type TextDocumentItem struct {
	URI     string `json:"uri"`
	Version int    `json:"version"`
	Text    string `json:"text"`
}

type TextDocumentIdentifier struct {
	URI string `json:"uri"`
}

type VersionedTextDocumentIdentifier struct {
	TextDocumentIdentifier
	Version int `json:"version"`
}

type NotificationDidOpen struct {
	jrpc2.Notification
	Params NotificationDidOpenParams `json:"params"`
}

type NotificationDidOpenParams struct {
	TextDocument TextDocumentItem `json:"textDocument"`
}

type NotificationDidChange struct {
	jrpc2.Notification
	Params NotificationDidChangeParams `json:"params"`
}

type NotificationDidChangeParams struct {
	TextDocument   VersionedTextDocumentIdentifier  `json:"textDocument"`
	ContentChanges []TextDocumentContentChangeEvent `json:"contentChanges"`
}

type TextDocumentContentChangeEvent struct {
	Text string `json:"text"`
}

type TextDocumentPositionParams struct {
	TextDocumentIdentifier `json:"textDocument"`
	Position               `json:"position"`
}

type Position struct {
	Line      uint `json:"line"`
	Character uint `json:"character"`
}

// Range represents a range in a text document expressed as (zero-based) start and
// end positions.
//
// https://microsoft.github.io/language-server-protocol/specifications/lsp/3.17/specification/#range
type Range struct {
	Start Position `json:"start"`
	End   Position `json:"end"`
}