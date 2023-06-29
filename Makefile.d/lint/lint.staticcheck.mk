lint/install/staticcheck:
	@(\
		command -v staticcheck &> /dev/null || { \
			echo "\033[34m>starting $@\033[0m"; \
			export STATICCHECK_CACHE="${{ runner.temp }}/staticcheck"; \
			go get -u honnef.co/go/tools; \
			go get honnef.co/go/tools/cmd/staticcheck; \
			go install "honnef.co/go/tools/cmd/staticcheck@2023.1.3"; \
			echo "Verify install..."; \
			command -v staticcheck || { \
				echo 'install failed.'; \
				exit 1; \
			}; \
			command -v staticcheck &> /dev/null || { \
				echo "\033[31m>(file not found) failed $@\033[0m";\
				ls -la /home/runner/.local/bin; \
				exit 1;\
			}; \
			echo "\033[32m>ok $@\033[0m";\
		};\
	)

lint/staticcheck/%:
	@echo "\033[34m>starting $@\033[0m"
	@GOOS=$(shell basename $@) staticcheck ./... || { \
		echo "\033[31m>failed $@\033[0m";\
		exit 1;\
	}; \
	echo "\033[32m>ok $@\033[0m"

lint/staticcheck: lint/install/staticcheck \
				  lint/staticcheck/darwin \
				  lint/staticcheck/linux \
				  lint/staticcheck/windows
	@echo "\033[32m>ok $@\033[0m"
