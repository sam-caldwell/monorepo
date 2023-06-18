#
# Makefile.d/_config/repo_root.mk
# (c) Sam Caldwell.  See LICENSE.txt
#
# This is the current working directory from Makefile's
# perspective, which should always be the root of our
# repo.
#
ifeq ($(OPSYS),windows)
REPO_ROOT:=$(shell echo %cd%)
else
REPO_ROOT:=$(shell pwd)
endif