expect/wget:
	command -v wget || { \
		echo "\033[31m>failed (wget missing) $@\033[0m";\
		exit 1;\
	}

security/snyk/install/darwin: expect/wget
	@echo "\033[34m>starting $@\033[0m"
	@rm -f snyk-macos
	@wget https://static.snyk.io/cli/latest/snyk-macos || { \
		echo "\033[31m>failed $@\033[0m";\
		exit 1;\
	}
	@mv ./snyk-macos $(shell go env GOPATH)/bin/snyk  || { \
		echo "\033[31m>failed $@\033[0m";\
		exit 1;\
	}
	@chmod a+x $(shell go env GOPATH)/bin/snyk  || { \
		echo "\033[31m>failed $@\033[0m";\
		exit 1;\
	}
	echo "\033[32m>ok $@\033[0m"

security/snyk/install/linux:
	@echo "\033[34m>starting $@\033[0m";\
	command -v wget || apt-get update -y && apt-get install wget -y || { \
		echo "\033[31m>failed $@\033[0m";\
		exit 1;\
	};\
	wget https://static.snyk.io/cli/latest/snyk-linux || { \
		echo "\033[31m>failed $@\033[0m";\
		exit 1;\
	};\
	mv snyk-linux $(shell go env GOPATH)/bin/snyk  || { \
		echo "\033[31m>failed $@\033[0m";\
		exit 1;\
	}



security/snyk/install:
	@command -v snyk &>/dev/null || {\
		echo "\033[34m>starting $@\033[0m";\
		make security/snyk/install/$(OPSYS);\
		echo "\033[32m>ok $@\033[0m";\
	}

