# make version/commit
# (c) 2023 Sam Caldwell.  See LICENSE.txt
#
# This file defines 'make version/commit' to add our updated
# version files and commit them to the git repo.
#
version/commit:
	@echo "\033[34m>starting $@\033[0m"
	@git add badges/VERSION.svg
	@git add projects/version/version.go
	@git commit -m "bump version to $(CURRENT_REPO_VERSION)" || { \
  		echo "\033[32m>nothing to commit\033[0m"; \
  		exit 1; \
    }; \
	@echo "\033[32m>ok $@\033[0m"