GO_VET_FLAGS:=-unsafeptr=false

lint/vet/%:
	@echo "\033[34m>starting $@\033[0m"
	@GOOS=$(shell basename $@) go vet $(GO_VET_FLAGS) ./... || { \
		echo "\033[31m>failed $@\033[0m";\
		exit 1;\
	}; \
	echo "\033[32m>ok $@\033[0m"

lint/vet: lint/vet/darwin \
          lint/vet/linux \
          lint/vet/windows
	@echo "\033[32m>ok $@\033[0m"
