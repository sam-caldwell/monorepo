#
# Makefile.d/_config/100.has_dpkg.package-manager.mk
# (c) Sam Caldwell.  See LICENSE.txt
#
# linux: debian/ubuntu apt package manager
#
HAS_DPKG=$(shell GOTMPDIR=$(GOTMPDIR) $(GO_BINARY) run cmd/tools/has-executable/main.go dpkg)
has_dpkg:
	@echo "$(HAS_DPKG)"