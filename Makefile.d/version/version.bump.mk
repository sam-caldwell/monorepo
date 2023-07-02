# make version/bump
# (c) 2023 Sam Caldwell.  See LICENSE.txt
#
# This file defines 'make version/bump' to detect the current version using
# the latest git tag on main branch, increment the patch version and update
# the tag for the local repository.
#
version/bump:
	@echo "\033[34m>starting $@\033[0m"
	@go run cmd/tools/bump-version/main.go -patch -updateTag
	@echo "\033[32m>ok $@\033[0m"