#
# PACKAGE MANAGERS
#
# windows:
HAS_CHOCO=$(shell go run cmd/tools/hasExecutable/main.go choco)
has_choco:
	@echo "$(HAS_CHOCO)"