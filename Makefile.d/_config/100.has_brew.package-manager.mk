#
# Makefile.d/_config/100.has_brew.package-manager.mk
# (c) Sam Caldwell.  See LICENSE.txt
#
# mac/linux package managers
#
HAS_BREW=$(shell GOTMPDIR=$(GOTMPDIR) $(GO_BINARY) run cmd/tools/hasExecutable/main.go brew)
has_brew:
	@echo "$(HAS_BREW)"
