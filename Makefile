#
# Makefile
# (c) 2023 Sam Caldwell.  See License.txt
#
# This is the root of our Makefile system.  It includes our Makefiles from Makefiles.d/
# and our Makefile tests in Makefiles.d/self-tests

include Makefile.d/*.mk
build:
	@go run cmd/tools/run-builds/main.go -color

clean:
	@go run cmd/tools/run-clean/main.go -color

hooks:
	@go run cmd/tools/update-git-hooks/main.go

lint:
	@go run cmd/tools/run-linter/main.go -color -quiet

security:
	@go run cmd/tools/run-scans/main.go

test:
	@go run cmd/tools/run-tests/main.go -color

version/bump:
	@go run cmd/tools/bump-version/main.go -patch -updateTag

version/commit:
	@git add badges/VERSION.svg
	@git add projects/version/version.go
	@git commit -m "bump version to $(CURRENT_REPO_VERSION)" || { \
  		echo "\033[32m>nothing to commit\033[0m"; \
  		exit 1; \
    }

version/badge:
	@go run cmd/tools/badge-maker/main.go -name VERSION -status "$(CURRENT_REPO_VERSION)" -color blue
	@go run cmd/tools/set-version/main.go