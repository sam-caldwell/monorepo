#
# PACKAGE MANAGERS
#
# linux: debian/ubuntu
HAS_APT=$(shell go run cmd/tools/hasExecutable/main.go apt-get)
has_apt:
	@echo "$(HAS_APT)"