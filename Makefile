.PHONY: all build test deps deps-cleancache

GOCMD=go
BUILD_DIR=build
BINARY_DIR=$(BUILD_DIR)/bin
CODE_COVERAGE=code-coverage

all: test build

build:
	go build -o bin/main main.go

wire:
	cd pkg/di && wire

run:
	go run main.go

build: ${BINARY_DIR}
	$(GOCMD) build -o $(BINARY_DIR) -v ./cmd/api

deps:
	$(GOCMD) get -u -t -d -v ./...
	$(GOCMD) mod tidy
	$(GOCMD) mod vendor

up:
	@docker compose up -d
