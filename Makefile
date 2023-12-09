.PHONY: dev-config build

build:
	go build -v ./cmd/server

dev-config:
	cp ./config/server.toml.local ./config/server.toml

test:
	go test -v -race ./...