#
# Makefile.d/lint/lint.install.lint.shell.mk
# (c) 2023 Sam Caldwell.  See LICENSE.txt
#
# Install the shellcheck linter.
#
lint/install/shellcheck/darwin:
	@echo "\033[34m>starting $@\033[0m"
	brew install shellcheck
	@echo "\033[32m>ok $@\033[0m"

lint/install/shellcheck/linux:
	@echo "\033[34m>starting $@\033[0m"
	apt-get install shellcheck
	@echo "\033[32m>ok $@\033[0m"

lint/install/shellcheck/windows:
	@echo "\033[31m> not implemented $@\033[0m"
	@exit 1

lint/install/shellcheck:
	@echo "\033[34m>starting $@\033[0m"
	@make lint/install/shellcheck/$(OPSYS)
	@echo "\033[32m>ok $@\033[0m"
