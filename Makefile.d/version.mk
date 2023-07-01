version:
	@go run cmd/bumpVersion/main.go -patch -updateTag
	@git push origin $(shell go run cmd/tools/bump-version/main.go)
