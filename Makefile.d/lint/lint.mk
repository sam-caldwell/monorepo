#
# Makefile.d/lint/lint.mk
# (c) 2023 Sam Caldwell.  See LICENSE.txt
#
# Run all the linters for the project against
# the entire repo.
#
lint:
	@go run cmd/tools/lint-repo/main.go -color -quiet
