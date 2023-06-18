clean:
	@echo "\033[32m>starting $@\033[0m"
	@make test/Makefile/teardown.fakeProject &> /dev/null
	@rm -rf ./bin || { echo -e "\033[31m>>FAILED:$@\033[0m";exit 1; }
	@mkdir -p ./bin || { echo -e "\033[31m>>FAILED:$@\033[0m";exit 1; }
	@rm -rf ./build || { echo -e "\033[31m>>FAILED:$@\033[0m";exit 1; }
	@mkdir -p ./build || { echo -e "\033[31m>>FAILED:$@\033[0m";exit 1; }
	@echo "\033[32m>clean...done\033[0m"
