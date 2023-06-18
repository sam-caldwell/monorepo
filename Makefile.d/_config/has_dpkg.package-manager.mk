#
# PACKAGE MANAGERS
#
# linux: debian/ubuntu
HAS_DPKG=$(shell go run cmd/tools/hasExecutable/main.go dpkg)
has_dpkg:
	@echo "$(HAS_DPKG)"