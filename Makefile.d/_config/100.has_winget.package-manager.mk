#
# Makefile.d/_config/100.has_winget.package-manager.mk
# (c) Sam Caldwell.  See LICENSE.txt
#
# linux: windows winget package manager
#
HAS_WINGET=$(shell GOTMPDIR=$(GOTMPDIR) $(GO_BINARY) run cmd/tools/hasExecutable/main.go winget)
has_winget:
	@echo "$(HAS_WINGET)"
