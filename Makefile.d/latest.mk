latest:
	@make latest/main

latest/%:
	git pull origin $(shell basename $@)