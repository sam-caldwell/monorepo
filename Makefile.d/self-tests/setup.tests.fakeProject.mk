#
#test/Makefile/setup.fakeProject.mk
# (c) 2023 Sam Caldwell.  See License.txt
#
test/Makefile/setup.fakeProject:
	@echo "\033[32m>starting $@\033[0m"
	@mkdir -p $(FAKE_COMMAND_DIR) || echo -e "\033[31m>FAILED:$@\033[0m"
	@mkdir -p $(FAKE_PROJECT_DIR) || echo -e "\033[31m>FAILED:$@\033[0m"
	@echo "\033[32m>done:$@\033[0m"
