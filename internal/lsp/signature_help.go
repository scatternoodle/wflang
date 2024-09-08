package lsp

import "github.com/scatternoodle/wflang/internal/jrpc2"

// Signature Help
//
// https://microsoft.github.io/language-server-protocol/specifications/lsp/3.17/specification/#textDocument_signatureHelp

// SignatureHelpOptions
//
// https://microsoft.github.io/language-server-protocol/specifications/lsp/3.17/specification/#signatureHelpOptions
type SignatureHelpOptions struct {
	TriggerChars   []string `json:"triggerCharacters,omitempty"`
	RetriggerChars []string `json:"retriggerCharacters,omitempty"`
}

// SignatureHelpRequest is sent from the client to the server to request signature information at a given cursor position.
//
// https://microsoft.github.io/language-server-protocol/specifications/lsp/3.17/specification/#textDocument_signatureHelp
type SignatureHelpRequest struct {
	jrpc2.Request
	SignatureHelpParams `json:"params"`
}

// SignatureHelpParams
//
// https://microsoft.github.io/language-server-protocol/specifications/lsp/3.17/specification/#signatureHelpParams
type SignatureHelpParams struct {
	TextDocumentPositionParams                       // Specifically, the position of the cursor at the point of request.
	Context                    *SignatureHelpContext `json:"context,omitempty"`
}

// SignatureHelpContext contains additional information about the context in which
// a signature help request was triggered.
//
// https://microsoft.github.io/language-server-protocol/specifications/lsp/3.17/specification/#signatureHelpContext
type SignatureHelpContext struct {
	TriggerKind         SignatureHelpTriggerKind `json:"triggerKind"`
	TriggerCharacter    string                   `json:"triggerCharacter,omitempty"`
	IsRetrigger         bool                     `json:"isRetrigger"`
	ActiveSignatureHelp *SignatureHelp           `json:"activeSignatureHelp,omitempty"`
}

// TriggerKind
//
// https://microsoft.github.io/language-server-protocol/specifications/lsp/3.17/specification/#signatureHelpTriggerKind
type SignatureHelpTriggerKind int

const (
	_ SignatureHelpTriggerKind = iota
	SighelpKindInvoked
	SighelpKindTriggerChar
	SighelpKindContentChange
)

// SignatureHelp represents the signature of a callable
//
// https://microsoft.github.io/language-server-protocol/specifications/lsp/3.17/specification/#signatureHelp
type SignatureHelp struct {
	Signatures      []SignatureInfo `json:"signatures"`
	ActiveSignature int             `json:"activeSignature,omitempty"`
	ActiveParameter int             `json:"activeParameter,omitempty"`
}

// SignatureInfo represents a parameter of a callable. A parameter can have a label
// and a doc-comment.
//
// https://microsoft.github.io/language-server-protocol/specifications/lsp/3.17/specification/#signatureInformation
type SignatureInfo struct {
	Label         string         `json:"label"`
	Documentation *MarkupContent `json:"documentation,omitempty"`
	Params        []ParamInfo    `json:"parameters,omitempty"`
	ActiveParam   int            `json:"activeparameter,omitempty"`
}

// ParamInfo represents a parameter of a callable-signature.
//
// https://microsoft.github.io/language-server-protocol/specifications/lsp/3.17/specification/#parameterInformation
type ParamInfo struct {
	// array of [start, end] inclusive start and end offsets of the parameter within the signature label.
	Label         [2]int         `json:"label"`
	Documentation *MarkupContent `json:"documentation,omitempty"`
}

// SignatureHelpResponse
//
// https://microsoft.github.io/language-server-protocol/specifications/lsp/3.17/specification/#textDocument_signatureHelp
type SignatureHelpResponse struct {
	jrpc2.Response
	*SignatureHelp `json:"result"`
}
