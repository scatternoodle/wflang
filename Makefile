build_vscode:
	go build -o editors/vscode/extension/server/bin/wflang cmd/server/main.go

build_repl:
	go build -o repl cmd/repl/main.go

build:
	make build_vscode
	make build_repl

test:
	go test ./...

run_repl:
	go run cmd/repl/main.go
