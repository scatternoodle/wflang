package lsp

import "github.com/scatternoodle/wflang/internal/jrpc2"

// ShowMessageNotification is sent from a server to a client to ask the client
// to display a particular message in the user interface
//
// https://microsoft.github.io/language-server-protocol/specifications/lsp/3.17/specification/#window_showMessage
type ShowMessageNotification struct {
	jrpc2.Notification
	Params ShowMessageParams `json:"params"`
}

// ShowMessageParams - see ShowMessageNotification
//
// https://microsoft.github.io/language-server-protocol/specifications/lsp/3.17/specification/#window_showMessage
type ShowMessageParams struct {
	Type    MessageType `json:"type"`
	Message string      `json:"message"`
}

// MessageType for window notifications
//
// https://microsoft.github.io/language-server-protocol/specifications/lsp/3.17/specification/#window_showMessage
type MessageType int

const (
	_ MessageType = iota
	Error
	Warning
	Info
	Log
	Debug
)
