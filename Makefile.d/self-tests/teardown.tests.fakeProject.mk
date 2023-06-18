#
#test/Makefile/teardown.fakeProject.mk
# (c) 2023 Sam Caldwell.  See License.txt
#
test/Makefile/teardown.fakeProject:
	@echo "starting $@"
	@$(RM_RF) cmd/fakeProject $(IGNORE_ERROR)
	@$(RM_RF)projects/fakeProject $(IGNORE_ERROR)
	@echo "done: $@"