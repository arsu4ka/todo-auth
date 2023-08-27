.PHONY: build
build:
	go build -o bin/server ./cmd/server/main.go

.PHONY: run
run:
	go run ./cmd/server/main.go

.PHONY: start
start:
	go build -o bin/server ./cmd/server/main.go && bin/server

.PHONY: test
test:
	go test -v -race ./...

.DEFAULT_GOAL := build
