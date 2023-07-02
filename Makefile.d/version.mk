CURRENT_REPO_VERSION=$(shell go run cmd/tools/bump-version/main.go)
version:
	@go run cmd/tools/bump-version/main.go -patch -updateTag
	@go run cmd/tools/badge-maker/main.go -name VERSION -status "$(CURRENT_REPO_VERSION)" -color blue
	@git add badges/VERSION.svg
	@git commit -m "bump version to $(CURRENT_REPO_VERSION)"
	@git push origin $(CURRENT_REPO_VERSION)
