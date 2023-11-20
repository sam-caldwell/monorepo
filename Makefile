include go/**/*.mk

clean:
	@echo "$@"
	@rm -rf build || true
	@mkdir build || true

build: clean
	@(\
		echo "$@"; \
		cd build && \
		echo "run child"; \
		make build/go/tools/numberCpuCores; \
	)

test:
	go test -v -failfast ./... && \
	go vet ./...

vet:
	go vet ./...
