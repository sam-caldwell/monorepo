lint/go/setup:
	@echo "\033[34m>starting $@\033[0m"
	@command -v staticcheck &> /dev/null || go install honnef.co/go/tools/cmd/staticcheck@latest || { \
		echo "\033[31m>failed $@\033[0m";\
		exit 1;\
	}; \
	echo "\033[32m>ok $@\033[0m"