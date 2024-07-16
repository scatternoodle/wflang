package lsp

import "github.com/scatternoodle/wflang/jrpc2"

type HoverRequest struct {
	jrpc2.Request
	TextDocumentPositionParams `json:"params"`
}

type HoverResponse struct {
	jrpc2.Response
	Hover `json:"result"`
}

type Hover struct {
	MarkupContent `json:"content"`
	Position      `json:"position,omitempty"`
}

type MarkupKind string

const (
	MarkupKindPlainText MarkupKind = "plaintext"
	MarkupKindMarkdown  MarkupKind = "markdown"
)

type MarkupContent struct {
	Kind  MarkupKind `json:"kind"`
	Value string     `json:"value"`
}
