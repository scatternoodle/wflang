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
