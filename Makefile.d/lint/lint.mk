# ToDo: Lint Makefile
# ToDo: Lint Markdown files
# ToDo: Lint Shell files

lint: lint/go \
      lint/yaml
	@echo "$@ done"
