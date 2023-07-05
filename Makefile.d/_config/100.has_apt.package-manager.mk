#
# PACKAGE MANAGERS
#
# linux: debian/ubuntu
HAS_APT=$(shell GOTMPDIR=$(GOTMPDIR) $(GO_BINARY) run cmd/tools/hasExecutable/main.go apt-get)

has_apt:
	@echo "$(HAS_APT)"