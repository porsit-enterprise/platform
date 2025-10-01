#╔═════════════════════════════════════════════════════════════════════════════════════════════════╗
#║ Copyright (C) 2025 porsit.com                                                                   ║
#╚═════════════════════════════════════════════════════════════════════════════════════════════════╝

.DEFAULT_GOAL := help

# --------------------------------------------------------------------------------------------------

help:
	@sed -n 's/^##//p' ${MAKEFILE_LIST} | column -t -s ':' |  sed -e 's/^/ /'
.PHONY: help

# --------------------------------------------------------------------------------------------------

## tag: Get latest tag
tag:
	@last_tag=$$(git describe --tags --abbrev=0 2>/dev/null); \
	if [ -z "$$last_tag" ]; then \
	  echo "No version tag found in the repository."; \
	  exit 1; \
	fi; \
	echo "Last tag: $$last_tag"; \
	new_version=$$(echo $$last_tag | awk -F. '{ $$3 = $$3 + 1; printf "%d.%d.%d", $$1, $$2, $$3 }'); \
	echo "New tag: $$new_version"
.PHONY: tag

## publish VERSION=<0.0.0>: Push application to Git
publish:
	@if [ -z "$$VERSION" ]; then echo "Error: VERSION environment variable is not set"; exit 1; fi
	git tag -a "v$(VERSION)" -m "Release v$(VERSION)"
	git push origin "v$(VERSION)"
.PHONY: publish

# --------------------------------------------------------------------------------------------------

## code.security: Inspects security problems by scanning the Go AST and SSA code representation
code.security:
	gosec -exclude-dir specifications -exclude-dir _principles ./...
.PHONY: code.security

## code.audit: Conduct quality checks
code.audit:
	go mod verify
	go vet ./...
	govulncheck ./...
.PHONY: code.audit

## code.lint: Run lint checks
code.lint:
	golangci-lint run
.PHONY: code.lint

## code.test: Run unit testes with coverage
test: code.security
	go test -v -timeout 30s -coverprofile=cover.out -cover ./...
	go tool cover -func=cover.out
.PHONY: code.test
