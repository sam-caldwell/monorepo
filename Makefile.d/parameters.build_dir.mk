
BUILD_PROJECTS := $(shell find cmd -mindepth 2 -maxdepth 3 -type d -execdir test ! -f "build.disabled" \; -print)
BUILD_DIR := ./build/
