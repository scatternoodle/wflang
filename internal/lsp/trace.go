package lsp

import "github.com/scatternoodle/wflang/internal/jrpc2"

// trace logging types
//
// https://microsoft.github.io/language-server-protocol/specifications/lsp/3.17/specification/#setTrace
// https://microsoft.github.io/language-server-protocol/specifications/lsp/3.17/specification/#logTrace

// SetTraceNotification is sent by the client to modify the trace setting of
// the server.
//
// https://microsoft.github.io/language-server-protocol/specifications/lsp/3.17/specification/#logTrace
type SetTraceNotification struct {
	jrpc2.Notification
	SetTraceParams `json:"params"`
}

// SetTraceParams
//
// https://microsoft.github.io/language-server-protocol/specifications/lsp/3.17/specification/#setTrace
type SetTraceParams struct {
	TraceValue `json:"value"`
}

type TraceValue string

const (
	TraceOff      TraceValue = "off"
	TraceMessages TraceValue = "messages"
	TraceVerbose  TraceValue = "verbose"
)

// LogTraceNotification sends a trace log of the server's execution to the client.
// Amount and content of the notification is controlled by the server's trace
// setting, set by LogTraceNotification from the client.
//
// https://microsoft.github.io/language-server-protocol/specifications/lsp/3.17/specification/#logTrace
type LogTraceNotification struct {
	jrpc2.Notification
	LogTraceParams `json:"params"`
}

// LogTraceParams
//
// https://microsoft.github.io/language-server-protocol/specifications/lsp/3.17/specification/#logTrace
type LogTraceParams struct {
	Message string `json:"message"`
	Verbose string `json:"verbose,omitempty"`
}
