#
# PACKAGE MANAGERS
#
# linux: redhat/centos/fedora/sles
HAS_RPM=$(shell go run cmd/tools/hasExecutable/main.go rpm)
has_rpm:
	@echo "$(HAS_RPM)"
