build:
	go build -o out cmd/main.go

build_vscode:
	go build -o editors/vscode/extension/server/bin/wflang cmd/main.go

test:
	go test ./...

run:
	go run cmd/main.go editors/vscode/extension/server/logs/server.log