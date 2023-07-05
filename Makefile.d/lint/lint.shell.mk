lint/shell: lint/install/shellcheck
	@echo "\033[34m>starting $@\033[0m"
	@shellcheck Makefile.d/scripts/*/*
	@find . -type f -name "*.sh" -print0 | while IFS= read -r -d $$'\0' file; do \
	    echo "\033[34m>>Checking $$file\033[0m"; \
	    shellcheck "$$file"; \
	    status=$$?; \
	    if [ $$status -ne 0 ]; then \
			echo "\033[32m>failed on $$file $@\033[0m"; \
	        exit $$status; \
	    fi; \
	done
	@echo "\033[32m>ok $@\033[0m"
