GOLANGCI_LINT = go tool -modfile tools/go.mod github.com/golangci/golangci-lint/v2/cmd/golangci-lint

.PHONY: lint
lint: golangci-lint

.PHONY: lint-fix
lint-fix: golangci-lint-fix

.PHONY: golangci-lint
golangci-lint: ## Run golangci-lint over the codebase.
	${GOLANGCI_LINT} run ./... --timeout 5m -v ${GOLANGCI_LINT_EXTRA_ARGS}

.PHONY: golangci-lint-fix
golangci-lint-fix: GOLANGCI_LINT_EXTRA_ARGS := --fix
golangci-lint-fix: golangci-lint ## Run golangci-lint over the codebase and run auto-fixers if supported by the linter
