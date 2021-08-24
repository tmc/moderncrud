# This file contains useful targets for development and testing purposes

.PHONY: test
test: deps ## Run tests
	go test -v ./...

.PHONY: generate
generate: ## Run code generation
	go generate ./...

.PHONY: deps
deps: ## Install dependencies.
	@command -v go > /dev/null || brew install go
	@command -v staticcheck > /dev/null || go get honnef.co/go/tools/cmd/staticcheck

.PHONY: lint
lint: lint-go ## Run linting.

.PHONY: lint-go
lint-go: deps ## Run linting for go code.
	@staticcheck ./...

