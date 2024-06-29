package lsp

const (
	// LifeCycle Messages

	MethodInitialize  = "initialize"
	MethodInitialized = "initialized"
	MethodShutdown    = "shutdown"
	MethodExit        = "exit"

	// Semantic Token Requests

	MethodSemanticTokensFull = "textDocument/semanticTokens/full"
)
