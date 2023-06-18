#
# Makefile.d/_config/build_projects.mk
# (c) 2023 Sam Caldwell.  See LICENSE.txt
#
# BUILD_PROJECTS is the list of projects in cmd/
#
ifeq ($(OPSYS),windows)
BUILD_PROJECTS := $(shell powershell -NoProfile -Command "Get-ChildItem -Path 'cmd' -Recurse -Directory | ForEach-Object { if (-not (Test-Path -Path (Join-Path $_.FullName 'build.disabled'))) { $_.FullName } }")
else
BUILD_PROJECTS := $(shell find cmd -mindepth 2 -maxdepth 3 -type d -execdir test ! -f "build.disabled" \; -print)
endif
