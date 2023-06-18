#
#test/Makefile/setup.fakeProject.mk
# (c) 2023 Sam Caldwell.  See License.txt
#
test/Makefile/setup.fakeProject:
	@echo "starting $@"
	@$(MKDIR) -p $(FAKE_COMMAND_DIR) || { echo "FAILED:$@";exit 1; }
	@$(MKDIR) -p $(FAKE_PROJECT_DIR) || { echo "FAILED:$@";exit 1; }
	@echo "done: $@"
