#
# PACKAGE MANAGERS
#
# windows:
HAS_WINGET=$(shell go run cmd/tools/hasExecutable/main.go winget)
has_winget:
	@echo "$(HAS_WINGET)"
