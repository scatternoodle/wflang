build_vscode:
	go build -o editors/vscode/extension/server/bin/wflang cmd/server/main.go

build_repl:
	go build -o repl cmd/repl/main.go

test:
	go test ./...

run_repl:
	go run cmd/repl/main.go