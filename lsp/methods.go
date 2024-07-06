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
	MethodDocDidSave   = "textDocument/didSave"

	// Semantic Token Requests

	MethodSemanticTokensFull = "textDocument/semanticTokens/full"
)
