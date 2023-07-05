#
# Makefile.d/_config/100.has_choco.package-manager.mk
# (c) Sam Caldwell.  See LICENSE.txt
#
# windows chocolatey package manager
#
HAS_CHOCO=$(shell GOTMPDIR=$(GOTMPDIR) $(GO_BINARY) run cmd/tools/hasExecutable/main.go choco)
has_choco:
	@echo "$(HAS_CHOCO)"