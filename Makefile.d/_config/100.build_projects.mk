#
# Makefile.d/_config/100.build_projects.mk
# (c) 2023 Sam Caldwell.  See LICENSE.txt
#
# BUILD_PROJECTS is the list of projects in cmd/
#
BUILD_PROJECTS := $(shell GOTMPDIR=$(GOTMPDIR) $(GO_BINARY) run cmd/tools/findBuildProjects/main.go cmd)


