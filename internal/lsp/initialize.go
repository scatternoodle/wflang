package lsp

// https://microsoft.github.io/language-server-protocol/specifications/lsp/3.17/specification/#lifeCycleMessages

import "github.com/scatternoodle/wflang/internal/jrpc2"

// InitializeRequest is always the first request sent by the client.
type InitializeRequest struct {
	jrpc2.Request
	Params InitializeRequestParams `json:"params"`
}

type InitializeRequestParams struct {
	// Parent process. Can be null if client process did not start the server
	ProcessID  *int     `json:"processId"`
	ClientInfo *AppInfo `json:"clientInfo,omitempty"`
	Locale     string   `json:"locale,omitempty"`

	// Deprecated in favour of rootURI
	RootPath *string `json:"rootPath,omitempty"`
	// Path of context root folder in the editor. Null if no folder context (i.e.
	// a single file was opened directly, or the app was started with no folder
	// selected).
	RootURI *string `json:"rootUri"`

	// User provided init options(?) TODO - what does this mean in practice?
	InitializationOptions ClientInitializationOptions `json:"initializationOptions,omitempty"`
	Capabilities          ClientCapabilities          `json:"capabilities"`

	// execution trace verbosity. Permitted values:
	//   "off" | "messages" | "verbose"
	Trace TraceValue `json:"trace"`
}

type ClientInitializationOptions struct {
	LogLevel string `json:"logLevel,omitempty"`
}

type AppInfo struct {
	Name    string `json:"name"`
	Version string `json:"version,omitempty"`
}

type InitializeResponse struct {
	jrpc2.Response
	Result InitializeResult `json:"result"`
}

type InitializeResult struct {
	Capabilities ServerCapabilities `json:"capabilities"`
	ServerInfo   *AppInfo           `json:"serverInfo,omitempty"`
}
