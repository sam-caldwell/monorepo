#
# Makefile.d/_config/100.has_dpkg.package-manager.mk
# (c) Sam Caldwell.  See LICENSE.txt
#
# linux: redhat/centos/fedora/sles rpm package manager
#
HAS_RPM=$(shell GOTMPDIR=$(GOTMPDIR) $(GO_BINARY) run cmd/tools/hasExecutable/main.go rpm)
has_rpm:
	@echo "$(HAS_RPM)"
