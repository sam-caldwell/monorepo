lint/setup:
	@go install golang.org/x/tools/go/analysis/passes/shadow/cmd/shadow@latest
	@command -v staticcheck &>/dev/null || go install honnef.co/go/tools/cmd/staticcheck@latest

lint/vet:
	@go vet ./...

lint: lint/setup \
      lint/vet
	staticcheck ./...
	@echo "$@ done"
