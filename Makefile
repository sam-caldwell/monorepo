# Makefile
# (c) 2023 Sam Caldwell.  See License.txt
#

include */Makefile.d/*.mk
include databases/*/Makefile.d/*.mk
#include containers/Makefile.d/*.mk
#include cpp/Makefile.d/*.mk
#include go/Makefile.d/*.mk
#include js/Makefile.d/*.mk
#include python/Makefile.d/*.mk
#include sql/Makefile.d/*.mk

clean: clean/asm \
	   clean/containers \
       clean/cpp \
	   clean/go \
	   clean/js \
	   clean/python \
	   clean/sql
	@rm -rf build || true
	@mkdir build || true
	color -green -lf ">>>Completed $@"
#
# Builds should--
#		- Build language-specific projects first.
#		- Build containers second (which may consume language-specific projects)
#		- Build the various database schemas next (which uses containers)
#
build: build/asm \
       build/cpp \
	   build/go \
	   build/js \
	   build/python \
       build/containers \
	   build/sql
	@color -yellow -lf ">>>Completed $@"

test:
	color -blue -lf ">>>Start $@"
	go test -v -failfast ./... && \
	go vet ./...
	color -green -lf "<<<Completed $@"

vet:
	color -green -lf ">>>Start $@"
	go vet ./...
	color -green -lf "<<<Completed $@"
