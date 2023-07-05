#
# Makefile.d/lint/lint.install.staticcheck.go.mk
# (c) 2023 Sam Caldwell.  See LICENSE.txt
#
# Check to ensure staticcheck and its dependencies are
# install and up to date and install them if not.
#

lint/install/staticcheck/tools:
	@GOTMPDIR=$(GOTMPDIR) $(GO_BINARY)  get -u "honnef.co/go/tools" || { \
		echo "\033[31m>install failed.  $@\033[0m"; \
		exit 1; \
	}

lint/install/staticcheck/get:
	@GOTMPDIR=$(GOTMPDIR) $(GO_BINARY)  get "honnef.co/go/tools/cmd/staticcheck" || { \
		echo "\033[31m>install failed.  $@\033[0m"; \
		exit 1; \
	};\

lint/install/staticcheck/install:
	@GOTMPDIR=$(GOTMPDIR) $(GO_BINARY) install "honnef.co/go/tools/cmd/staticcheck@latest" || { \
		echo "\033[31m>install failed.  $@\033[0m"; \
		exit 1; \
	};\

#export STATICCHECK_CACHE="${{ runner.temp }}/staticcheck"
lint/install/staticcheck:
	@echo "\033[34m>starting $@\033[0m"; \
	command -v staticcheck &>/dev/null || { \
		make -s lint/install/staticcheck/tools;\
		make -s lint/install/staticcheck/get;\
	};\
	make -s lint/install/staticcheck/install; \
	command -v staticcheck &>/dev/null || {\
		echo "\033[31m>install failed.  $@\033[0m"; \
		exit 1; \
  	};\
	echo "\033[32m>ok $@\033[0m"
