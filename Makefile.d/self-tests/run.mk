#
# self-tests/run.mk
# (c) 2023 Sam Caldwell.  See License.txt
#

.PHONY: test/self-tests/run
test/self-tests/run: clean
	@echo "\033[32m>starting $@\033[0m"
	@echo "\033[33m>>WARNING: test/self-tests/run functionality is disabled/not implemented.\033[0m"
	#@make test/Makefile/disable.test  || { echo -e "\033[31m>>FAILED:$@\033[0m";exit 1; }
	#@make test/Makefile/enable.test  || { echo -e "\033[31m>>FAILED:$@\033[0m";exit 1; }
	#ToDo: add more tests here.
	@echo "\033[32m>done:$@\033[0m"
