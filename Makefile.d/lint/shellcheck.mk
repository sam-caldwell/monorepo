lint/shell:
	@echo "\033[34m>starting $@\033[0m"
	@shellcheck Makefile.d/scripts/*/*
	@shellcheck **/.sh
	@echo "\033[32m>ok $@\033[0m"
