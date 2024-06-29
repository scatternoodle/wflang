package lsp

import "github.com/scatternoodle/wflang/jrpc2"

type SemanticTokensFullRequest struct {
	jrpc2.Request
	Params SemanticTokensFullParams `json:"params"`
}

type SemanticTokensFullParams struct {
	TextDocument TextDocumentIdentifier `json:"textDocument"`
}
