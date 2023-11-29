# Makefile
# (c) 2023 Sam Caldwell.  See License.txt
#

RED:="\033[31m"
GREEN:="\033[32m"
BLUE:="\033[34m"
YELLOW:="\033[33m"
RESET:="\033[0m"

include containers/Makefile.d/*.mk
#include go/**/*.mk

clean:
	@echo "$@"
	@rm -rf build || true
	@mkdir build || true

build: clean \
	   build/containers/
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
