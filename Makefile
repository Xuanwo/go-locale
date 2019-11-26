SHELL := /bin/bash

.PHONY: all check format vet lint build test tidy

help:
	@echo "Please use \`make <target>\` where <target> is one of"
	@echo "  check      to format, vet and lint "
	@echo "  build      to create bin directory and build"
	@echo "  generate   to generate code"
	@echo "  test       to run test"

# golint: go get -u golang.org/x/lint/golint
tools := golint

$(tools):
	@command -v $@ >/dev/null 2>&1 || echo "$@ is not found, plese install it."

check: format vet lint

format:
	@echo "go fmt"
	@go fmt ./...
	@echo "ok"

vet:
	@echo "go vet"
	@go vet ./...
	@echo "ok"

lint: golint
	@echo "golint"
	@golint ./...
	@echo "ok"

build: tidy check
	@echo "build go-language"
	@go build ./...
	@echo "ok"

test:
	@echo "run test"
	@go test -race -coverprofile=coverage.txt -covermode=atomic -v ./...
	@go tool cover -html="coverage.txt" -o "coverage.html"
	@echo "ok"

tidy:
	@echo "Tidy and check the go mod files"
	@go mod tidy && go mod verify
	@echo "Done"
