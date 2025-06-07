# Makefile for MCP-PoliticalCompass

# Variables
BINARY_NAME := mcp-political-compass
VERSION := $(shell cat VERSION 2>/dev/null || echo "dev")
BUILD_DIR := dist

# Default target
.PHONY: all
all: test build

# Build for current platform
.PHONY: build
build:
	go build -ldflags "-X main.Version=$(VERSION)" -o $(BINARY_NAME) .

# Run tests
.PHONY: test
test:
	go test -v ./...

# Run tests with coverage
.PHONY: test-coverage
test-coverage:
	go test -coverprofile=coverage.out ./...
	go tool cover -html=coverage.out -o coverage.html

# Cross-compile for all platforms
.PHONY: build-all
build-all:
	./build.sh $(VERSION)

# Clean build artifacts
.PHONY: clean
clean:
	rm -f $(BINARY_NAME)
	rm -rf $(BUILD_DIR)
	rm -f coverage.out coverage.html

# Install dependencies
.PHONY: deps
deps:
	go mod download
	go mod tidy

# Run linter
.PHONY: lint
lint:
	go vet ./...
	go fmt ./...

# Development server (rebuild on changes)
.PHONY: dev
dev: build
	./$(BINARY_NAME)

# Show version
.PHONY: version
version:
	@echo $(VERSION)

# Help
.PHONY: help
help:
	@echo "Available targets:"
	@echo "  build        - Build for current platform"
	@echo "  build-all    - Cross-compile for all platforms"
	@echo "  test         - Run tests"
	@echo "  test-coverage- Run tests with coverage report"
	@echo "  clean        - Clean build artifacts"
	@echo "  deps         - Install/update dependencies"
	@echo "  lint         - Run linter and formatter"
	@echo "  dev          - Build and run development server"
	@echo "  version      - Show current version"
	@echo "  help         - Show this help message"
