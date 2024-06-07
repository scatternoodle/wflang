build:
	go build -o out cmd/main.go

test:
	go test ./...

run:
	make build && ./out