package lsp

func serverCababilities() ServerCapabilities {
	return ServerCapabilities{

		SemanticTokensProvider: SemanticTokensOptions{},
	}
}

type ServerCapabilities struct {
	SemanticTokensProvider SemanticTokensOptions `json:"semanticTokensProvider,omitempty"`
}

type SemanticTokensOptions struct {
	Legend TokenTypesLegend `json:"legend"`
	Range  bool             `json:"range,omitempty"`
	Full   bool             `json:"full,omitempty"`
}

type TokenTypesLegend struct {
	TokenTypes     []string `json:"tokenTypes"`
	TokenModifiers []string `json:"tokenModifiers"`
}
