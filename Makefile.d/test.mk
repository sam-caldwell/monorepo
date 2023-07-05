test: lint test/simple
	@echo "\033[32m>done:$@\033[0m"

test/simple: # test/Makefile.d/self-tests/run
	@echo "\033[32m>starting $@\033[0m"
	go test -failfast -v -count=1 -vet=all ./... || { \
  		echo "\033[31m>failed:$@\033[0m"; \
  		exit 1; \
  	}; \
	echo "\033[32m>done:$@\033[0m"


tests:
	@echo "\033[31m>Did you mean 'make -s test'?\033[0m"
	@exit 1