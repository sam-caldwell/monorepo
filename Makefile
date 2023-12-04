# Makefile
# (c) 2023 Sam Caldwell.  See License.txt
#

RED:="\033[31m"
GREEN:="\033[32m"
BLUE:="\033[34m"
YELLOW:="\033[33m"
RESET:="\033[0m"

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
	@echo "$@"
	@rm -rf build || true
	@mkdir build || true
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
	@echo ${GREEN}">>>Completed $@"${RESET}
#
#	@(\
#		echo "$@"; \
#		cd build && \
#		echo "run child"; \
#		make build/go/tools/numberCpuCores; \
#	)

test:
	go test -v -failfast ./... && \
	go vet ./...

vet:
	go vet ./...
