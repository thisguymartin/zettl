.PHONY: dev

build:
	@go build -o bin/cli ./cmd/cli/main.go

dev:
	@go run ./cmd/cli/main.go && ./air.toml

run:
	@echo "Watching for changes..."
	@air -c ./.air.toml 