#
# Makefile.d/lint/lint.go.mk
# (c) 2023 Sam Caldwell.  See LICENSE.txt
#
# Perform go-related linting operations and handle the
# GOOS/GOARCH loops need to ensure broad coverage.
#
lint/go:
	@echo "\033[34m>starting $@\033[0m"
	@echo "\033[34m>> install prerequisites $@\033[0m"
	@make lint/install/staticcheck
	@echo "\033[32m>> install prerequisites OK\033[0m"
	@echo "\033[34m>starting linter cycle $@\033[0m"
#	@(\
#		for THIS_OS in $(GOOS_LIST); do \
#			for THIS_ARCH in $(GOARCH_LIST); do \
#				GOOS=$(THIS_OS) GOARCH=$(THIS_ARCH) $(GO_BINARY) vet $(GO_VET_FLAGS) ./... || { \
#					echo "\033[31m>failed [OP:'go vet' GOOS:'$(THIS_OS)' GOARCH:'$(THIS_ARCH)'] $@\033[0m";\
#					exit 1;\
#				}; \
#				GOOS=$(THIS_OS) GOARCH=$(THIS_ARCH) $(GO_BINARY) staticcheck ./... || { \
#					echo "\033[31m>failed [OP:'staticcheck' GOOS:'$(THIS_OS)' GOARCH:'$(THIS_ARCH)'] $@\033[0m";\
#					exit 1;\
#				}; \
#			done; \
#		done; \
#	) || {\
#	  echo "\033[31m>failed $@\033[0m";\
#	  exit 1;\
#  	};\
#  	echo "\033[32m>ok $@\033[0m"
