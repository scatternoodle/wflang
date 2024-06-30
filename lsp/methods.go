package lsp

const (
	// LifeCycle Messages

	MethodInitialize  = "initialize"
	MethodInitialized = "initialized"
	MethodShutdown    = "shutdown"
	MethodExit        = "exit"

	// Document Sync
	MethodDocDidOpen   = "textDocument/didOpen"
	MethodDocDidChange = "textDocument/didChange"

	// Semantic Token Requests

	MethodSemanticTokensFull = "textDocument/semanticTokens/full"
)
