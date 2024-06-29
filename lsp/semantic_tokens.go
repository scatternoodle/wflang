package lsp

import "github.com/scatternoodle/wflang/jrpc2"

type RequestSemanticTokensFull struct {
	jrpc2.Request
	Params SemanticTokensFullParams `json:"params"`
}

type SemanticTokensFullParams struct {
	TextDocument TextDocumentIdentifier `json:"textDocument"`
}

type ResponseSemanticTokensFull struct {
	jrpc2.Response
	Result SemanticTokens `json:"result"`
}

type SemanticTokens struct {
	ResultID *string `json:"resultId,omitempty"`
	Data     []Uint
}
