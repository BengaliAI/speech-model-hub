.SHELL=/usr/bin/bash
.PHONY: all clean fmt tidy

all: clean fmt tidy build

clean:
	@echo "Cleaning..."
	@rm -rf bin
	@rm -rf tmp/*

fmt:
	@echo "Formatting..."
	@go fmt ./...

tidy:
	@echo "Tidying..."
	@go mod tidy

build:
	@echo "Building..."
	@mkdir -p bin
	@go build -o bin/ ./...