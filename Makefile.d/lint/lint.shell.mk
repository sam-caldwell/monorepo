lint/shell: lint/install/shellcheck
	@echo "\033[34m>starting $@\033[0m"
	@shellcheck Makefile.d/scripts/*/*
	@( \
		for file in $(find . -type f -name "*.sh" -print0); do \
			shellcheck $${file} || {\
				echo "\033[32m>failed on $${file} $@\033[0m"; \
				exit $$status; \
			};\
	  	done;\
	) || exit 1;\
	echo "\033[32m>ok $@\033[0m"
