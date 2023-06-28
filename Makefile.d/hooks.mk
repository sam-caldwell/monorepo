hooks:
	@rm -rf .git/hooks &> /dev/null || true
	@cp -rfvp Makefile.d/scripts/hooks .git/
