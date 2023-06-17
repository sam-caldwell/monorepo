lint/setup:
	@go install golang.org/x/tools/go/analysis/passes/shadow/cmd/shadow@latest
	@go install honnef.co/go/tools/cmd/staticcheck@latest

lint/vet:
	@go vet ./...

lint/staticcheck:
	@staticcheck ./...

lint: lint/setup \
      lint/vet \
      lint/staticcheck
	@echo "$@ done"
