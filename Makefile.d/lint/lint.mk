# ToDo: Lint Makefile
# ToDo: Lint Markdown files
# ToDo: Lint Shell files

lint: lint/go \
      lint/yaml
	@echo "\033[32m>ok $@\033[0m"
