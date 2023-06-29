#
# test/Makefile/enable.test.mk
# (c) 2023 Sam Caldwell.  See License.txt
#
test/Makefile/enable.test:
	@echo "\033[32m>>starting $@\033[0m"
	@make test/Makefile/setup.fakeProject || echo -e "\033[31m>>FAILED:$@\033[0m"
	@make enable/tests/fakeProject || echo -e "\033[31m>>FAILED:$@\033[0m"
	@( \
		export PROJECT_DIR="./cmd/fakeProject/build.disabled"; \
		export CMD_FLAG_FILE="./cmd/fakeProject/build.disabled"; \
		if [ -f "$${CMD_FLAG_FILE}" ]; then \
			@echo "\033[31m>Error: Enable operation should have removed $${CMD_FLAG_FILE}:$@\033[0m"
	  	fi; \
	)
	@make test/Makefile.d/self-tests/teardown.fakeProject
	@echo "\033[32m>>done:$@\033[0m"
