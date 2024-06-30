package lsp

import "github.com/scatternoodle/wflang/jrpc2"

type TextDocumentIdentifier struct {
	URI string `json:"uri"`
}

type NotificationDidOpen struct {
	jrpc2.Notification
	Params NotificationDidOpenParams `json:"params"`
}

type NotificationDidOpenParams struct {
	TextDocument TextDocumentItem `json:"textDocument"`
}

type TextDocumentItem struct {
	URI     string `json"uri"`
	Version Int    `json:"version"`
	Text    string `json:"text"`
}
