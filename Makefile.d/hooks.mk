hooks:
	@rm -rf .git/hooks &> /dev/null || true
	@ln -sf Makefile.d/scripts/hooks .git/hooks
