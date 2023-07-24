#
# For build automation docs, see docs/builds/README.md
#
.PHONY: build
build:
	@go run cmd/tools/run-builds/main.go -color
