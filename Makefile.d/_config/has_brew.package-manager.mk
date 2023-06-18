#
# PACKAGE MANAGERS
#
# mac/linux package managers
HAS_BREW=$(shell go run cmd/tools/hasExecutable/main.go brew)
has_brew:
	@echo "$(HAS_BREW)"
