GOLANGCI_LINT = go tool -modfile tools/go.mod github.com/golangci/golangci-lint/v2/cmd/golangci-lint
LEFTHOOK = go tool -modfile tools/go.mod github.com/evilmartians/lefthook

# Default fuzz test duration
FUZZ_TIME ?= 30s

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

.PHONY: fmt
fmt: ## Format code with golangci-lint
	${GOLANGCI_LINT} fmt ./...

.PHONY: fmt-diff
fmt-diff: ## Show code formatting differences
	${GOLANGCI_LINT} fmt --diff ./...

# Install lefthook
.PHONY: install-lefthook
install-lefthook:
	${LEFTHOOK} install

# Documentation targets
.PHONY: docs-dev
docs-dev: ## Start Hugo development server
	cd docs && hugo server

# Install govalid binary
.PHONY: install-govalid
install-govalid: ## Install govalid binary for validation code generation
	go install ./cmd/govalid

# Generate validation code for test
.PHONY: generate-validation-code
generate-validation-code: install-govalid ## Generate validation code using govalid
	go generate ./test/marker.go

# Generate new validator scaffold
.PHONY: generate-validator
generate-validator: ## Generate a new validator scaffold and all registry files. Usage: make generate-validator MARKER=phoneNumber
	go run cmd/generate-validators/main.go -marker=$(MARKER)

# Test targets
.PHONY: test
test: ## Run all tests except validation helper (due to known issues)
	go test ./... -shuffle on -v -race
	go test -C test ./... -shuffle on -v -race
	go test ./... -shuffle on -v -race -tags=test

# Fuzz test targets
.PHONY: fuzz
fuzz: fuzz-email fuzz-uuid fuzz-url ## Run all fuzz tests

.PHONY: fuzz-email
fuzz-email: ## Run email validation fuzz test
	@echo "Running email validation fuzz test for $(FUZZ_TIME)..."
	cd validation/validationhelper && go test -run "^$$" -fuzz=FuzzIsValidEmail -fuzztime=$(FUZZ_TIME) -v

.PHONY: fuzz-uuid
fuzz-uuid: ## Run UUID validation fuzz test
	@echo "Running UUID validation fuzz test for $(FUZZ_TIME)..."
	cd validation/validationhelper && go test -run "^$$" -fuzz=FuzzIsValidUUID -fuzztime=$(FUZZ_TIME) -v

.PHONY: fuzz-url
fuzz-url: ## Run URL validation fuzz test
	@echo "Running URL validation fuzz test for $(FUZZ_TIME)..."
	cd validation/validationhelper && go test -run "^$$" -fuzz=FuzzIsValidURL -fuzztime=$(FUZZ_TIME) -v

.PHONY: fuzz-quick
fuzz-quick: ## Run quick fuzz tests (15 seconds each)
	@echo "Running quick fuzz tests (15 seconds each)..."
	$(MAKE) fuzz FUZZ_TIME=15s

.PHONY: fuzz-long
fuzz-long: ## Run long fuzz tests (5 minutes each)
	@echo "Running long fuzz tests (5 minutes each)..."
	$(MAKE) fuzz FUZZ_TIME=5m

.PHONY: fuzz-ci
fuzz-ci: ## Run fuzz tests for CI (30 seconds each)
	@echo "Running CI fuzz tests (30 seconds each)..."
	$(MAKE) fuzz FUZZ_TIME=30s

.PHONY: fuzz-dev
fuzz-dev: ## Run development fuzz tests (1 minute each)
	@echo "Running development fuzz tests (1 minute each)..."
	$(MAKE) fuzz FUZZ_TIME=1m

# Documentation targets
.PHONY: docs-serve
docs-serve: ## Serve documentation site locally
	@echo "Starting Hugo development server..."
	@command -v hugo >/dev/null 2>&1 || { echo "Hugo is not installed. Please install Hugo first: https://gohugo.io/installation/"; exit 1; }
	cd docs && hugo server -D

.PHONY: docs-build
docs-build: ## Build documentation site for production
	@echo "Building documentation site..."
	@command -v hugo >/dev/null 2>&1 || { echo "Hugo is not installed. Please install Hugo first: https://gohugo.io/installation/"; exit 1; }
	cd docs && hugo --minify

.PHONY: docs-install
docs-install: ## Install Hugo (macOS only)
	@echo "Installing Hugo..."
	@if command -v brew >/dev/null 2>&1; then \
		brew install hugo; \
	else \
		echo "Homebrew not found. Please install Hugo manually: https://gohugo.io/installation/"; \
		exit 1; \
	fi


.PHONY: sync-benchmarks
sync-benchmarks: ## Synchronize benchmark results across all documentation files
	@echo "Synchronizing benchmark results..."
	./scripts/sync-benchmarks.sh



.PHONY: help
help: ## Show this help message
	@echo 'Usage: make [target] [FUZZ_TIME=duration]'
	@echo ''
	@echo 'Targets:'
	@awk 'BEGIN {FS = ":.*?## "} /^[a-zA-Z_-]+:.*?## / {printf "  \033[36m%-15s\033[0m %s\n", $$1, $$2}' $(MAKEFILE_LIST)
	@echo ''
	@echo 'Examples:'
	@echo '  make fuzz-quick         # Run quick fuzz tests (15s each)'
	@echo '  make fuzz-email         # Run email fuzz test (30s default)'
	@echo '  make fuzz FUZZ_TIME=2m  # Run all fuzz tests for 2 minutes each'
	@echo '  make test               # Run regular tests (excluding validation helper)'
	@echo '  make lint               # Run linter'
	@echo '  make docs-serve         # Serve documentation site locally'
	@echo '  make docs-build         # Build documentation for production'
	@echo '  make sync-benchmarks    # Sync benchmark results across all docs'
