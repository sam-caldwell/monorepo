all: build-client build-server build-query-compile

lint:
	@echo 'linter not implemented yet'

test:
	#unit tests
	go test -v ./...

build-client: lint test
	go build -o ${BUILD_DIR}/osRecon-client main.client.go

build-server: lint test
	go build -o ${BUILD_DIR}/osRecon-server main.server.go

build-query-compile: lint test
	go build -o ${BUILD_DIR}/query-compile main.query-compile.go

integration-tests: build-client build-server build-query-compile
	@echo '$@ not implemented'
	@exit 1
