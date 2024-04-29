all: build

lint:
	go vet -v ./...

test:
	go test -v ./...

build: build-jira build-confluence build-statuspage

build-confluence: lint test
	@echo "$@not implemented"

build-jira: lint test
	@echo "$@not implemented"

build-statuspage: lint test
	@echo "$@not implemented"
