#
# self-tests/disable.tests.mk
# (c) 2023 Sam Caldwell.  See License.txt
#
test/Makefile/disable.test:
	@echo "\033[32m>>starting $@\033[0m"

	@make test/Makefile/setup.fakeProject || echo -e "\033[31m>>FAILED:$@\033[0m"
	@make enable/tests/fakeProject || echo -e "\033[31m>>FAILED:$@\033[0m"
	@( \
		export PROJECT_DIR="$${FAKE_PROJECT_DIR}/test.disabled"; \
		export CMD_FLAG_FILE="$${FAKE_COMMAND_DIR}/test.disabled"; \
		if [ !-f "$${CMD_FLAG_FILE}" ]; then \
			@echo "\033[31m>Error: Enable operation should have created $${CMD_FLAG_FILE}:$@\033[0m"
	  	fi; \
	)
	@make test/Makefile/teardown.fakeProject || echo -e "\033[31m>FAILED:$@\033[0m"
	@echo "\033[32m>>done:$@\033[0m"
