package lsp

import "github.com/scatternoodle/wflang/jrpc2"

type SemanticTokensOptions struct {
	Legend TokenTypesLegend `json:"legend"`
	Range  bool             `json:"range,omitempty"`
	Full   bool             `json:"full,omitempty"`
}

type TokenTypesLegend struct {
	TokenTypes     []string `json:"tokenTypes"`
	TokenModifiers []string `json:"tokenModifiers"`
}

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
