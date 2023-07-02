# make version/badge
# (c) 2023 Sam Caldwell.  See LICENSE.txt
#
# This file defines 'make version/badge' to update the 'badges/VERSION.svg'
# badge file with the current version number.  This will also run 'set-version'
# to update the version.go constant file.
#
version/badge:
	@echo "\033[34m>starting $@\033[0m"
	@go run cmd/tools/badge-maker/main.go -name VERSION -status "$(CURRENT_REPO_VERSION)" -color blue
	@go run cmd/tools/set-version/main.go
	@echo "\033[32m>ok $@\033[0m"
