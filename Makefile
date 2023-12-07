# Makefile
# (c) 2023 Sam Caldwell.  See License.txt
#
all:
	@go run go/tools/monorepo/main.go list clean build test

build:
	@go run go/tools/monorepo/main.go build

clean:
	@go run go/tools/monorepo/main.go clean

list:
	@go run go/tools/monorepo/main.go list

test:
	@go run go/tools/monorepo/main.go test
