#
# Makefile.d/lint/lint.mk
# (c) 2023 Sam Caldwell.  See LICENSE.txt
#
# Run all the linters for the project against
# the entire repo.
#
lint: lint/go \
      lint/yaml \
      lint/shell
	@echo "\033[32m>ok $@\033[0m"
