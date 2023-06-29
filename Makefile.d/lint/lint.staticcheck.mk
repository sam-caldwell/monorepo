lint/staticcheck/%:
	@echo "\033[34m>starting $@\033[0m"
	@GOOS=$(shell basename $@) staticcheck ./... || { \
		echo "\033[31m>failed $@\033[0m";\
		exit 1;\
	}; \
	echo "\033[32m>ok $@\033[0m"

lint/staticcheck: lint/staticcheck/darwin # \
#				  lint/staticcheck/linux \
#				  lint/staticcheck/windows
	echo "\033[32m>ok $@\033[0m"