hooks:
	@rm -rf .git/hooks &> /dev/null || true
	@mkdir -p .git/hooks
	@cp -rfvp Makefile.d/scripts/hooks .git/
