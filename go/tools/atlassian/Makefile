BUILD_DIR:=../../../build/atlassian

all: build

setup:
	@echo "initializing ${BUILD_DIR}"
	@rm -rf ${BUILD_DIR} || true
	@mkdir -p ${BUILD_DIR}

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
