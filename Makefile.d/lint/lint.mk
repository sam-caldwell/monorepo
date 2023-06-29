# ToDo: Lint Makefile
# ToDo: Lint Markdown files
# ToDo: Lint Yaml files

lint: lint/go \
      lint/yaml
	@echo "$@ done"
