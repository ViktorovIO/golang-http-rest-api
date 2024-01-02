ROOT_DIR:=$(shell dirname $(realpath $(firstword $(MAKEFILE_LIST))))

.PHONY: dev-config build

# Local machine

build:
	go build -v ./cmd/server

dev-config:
	cp ./config/server.toml.local ./config/server.toml

# Docker

up:
	docker compose up -d

run:
	docker compose exec -u www-data golang-server ./server

test:
	go test -v -race ./...