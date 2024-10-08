package lsp

const (
	MethodInitialize         string = "initialize"
	MethodInitialized        string = "initialized"
	MethodShutdown           string = "shutdown"
	MethodExit               string = "exit"
	MethodDocDidOpen         string = "textDocument/didOpen"
	MethodDocDidChange       string = "textDocument/didChange"
	MethodDocDidSave         string = "textDocument/didSave"
	MethodSemanticTokensFull string = "textDocument/semanticTokens/full"
	MethodHover              string = "textDocument/hover"
	MethodDocumentSymbols    string = "textDocument/documentSymbol"
	MethodDeclaration        string = "textDocument/declaration"
	MethodDefinition         string = "textDocument/definition"
	MethodShowMessage        string = "window/showMessage"
	MethodCompletion         string = "textDocument/completion"
	MethodRename             string = "textDocument/rename"
	MethodSignatureHelp      string = "textDocument/signatureHelp"
	MethodSetTrace           string = "$/setTrace"
	MethodLogTrace           string = "$/logTrace"
)
