#
# Makefile.d/_config/build_projects.mk
# (c) 2023 Sam Caldwell.  See LICENSE.txt
#
# BUILD_PROJECTS is the list of projects in cmd/
#
BUILD_PROJECTS := $(shell go run cmd/tools/findBuildProjects/main.go cmd)


