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

// Location represents a location inside a resource, such as a line inside a text file.
//
// https://microsoft.github.io/language-server-protocol/specifications/lsp/3.17/specification/#location
type Location struct {
	URI   string `json:"uri"`
	Range Range  `json:"range"`
}

// LocationLink represents a link between a source and a target location. Key
// difference between this and lsp.Location is that LocationLink supports mouse
// interaction (e.g. ctrl-clicking for a goto-definition in VSCode).
//
// https://microsoft.github.io/language-server-protocol/specifications/lsp/3.17/specification/#locationLink
type LocationLink struct {
	// Span of the origin of the link for mouseover. If omitted, defaults to word range at mouse position.
	OriginSelectionRange Range `json:"originSelectionRange,omitempty"`

	TargetURI string `json:"targetURI"`

	// Range encompassing where the range you want the cursor to be able to pick
	// up the symbol. Can be larger than the SelectionRange, which strictly
	// encompasses the symbol itself.
	TargetRange Range `json:"targetRange"`

	// SelectionRange encompasses the symbol itself, and is used to highlight the
	// symbol in the editor / drive selection.
	TargetSelectionRange Range `json:"targetSelectionRange"`
}
