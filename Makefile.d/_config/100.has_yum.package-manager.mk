#
# Makefile.d/_config/100.has_yum.package-manager.mk
# (c) Sam Caldwell.  See LICENSE.txt
#
# linux: redhat/centos/fedora yum
#
HAS_YUM=$(shell GOTMPDIR=$(GOTMPDIR) $(GO_BINARY) run cmd/tools/hasExecutable/main.go yum)
has_yum:
	@echo "$(HAS_YUM)"
