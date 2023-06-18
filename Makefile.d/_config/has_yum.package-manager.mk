#
# PACKAGE MANAGERS
#
# linux: redhat/centos/fedora
HAS_YUM=$(shell go run cmd/tools/hasExecutable/main.go yum)
has_yum:
	@echo "$(HAS_YUM)"
