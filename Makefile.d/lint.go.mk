

lint/go/setup/snyk: lint/go/setup/snyk/macos
	@echo "$@ done"

lint/go/setup:
	@echo "\033[34m>starting $@\033[0m"
	@command -v staticcheck &> /dev/null || go install honnef.co/go/tools/cmd/staticcheck@latest || { \
		echo "\033[31m>failed $@\033[0m";\
		exit 1;\
	}; \
	echo "\033[32m>ok $@\033[0m"

lint/vet:
	@echo "\033[34m>starting $@\033[0m"
	@(\
		GOOS=darwin go vet ./... || { \
			echo "\033[31m>failed on darwin [$@]\033[0m";\
			exit 1;\
		} || exit 1;\
		echo "\033[32m>darwin: ok $@\033[0m"; \
		GOOS=linux go vet ./... || { \
			echo "\033[31m>failed on linux [$@]\033[0m";\
			exit 1;\
		} || exit 1;\
		echo "\033[32m>linux: ok $@\033[0m"; \
		GOOS=windows go vet ./... || { \
			echo "\033[31m>failed on windows [$@]\033[0m";\
			exit 1;\
		} || exit 1; \
		echo "\033[32m>windows: ok $@\033[0m"; \
	);\
	echo "\033[32m>ok $@\033[0m"

lint/staticcheck:
	@echo "\033[34m>starting $@\033[0m"
	@staticcheck ./... || { \
		echo "\033[31m>failed $@\033[0m";\
		exit 1;\
	}; \
	echo "\033[32m>ok $@\033[0m"

lint/go: lint/go/setup \
		 lint/vet \
		 lint/staticcheck
	echo "\033[32m>ok $@\033[0m"