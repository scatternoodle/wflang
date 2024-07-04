package lsp

import "github.com/scatternoodle/wflang/jrpc2"

type SemanticTokensRequest struct {
	jrpc2.Request
	TextDocumentIdentifier `json:"textDocument"`
}

type SemanticTokensResponse struct {
	jrpc2.Response
	Result SemanticTokensResult `json:"result"`
}

type SemanticTokensResult struct {
	Data []uint `json:"data"`
}
