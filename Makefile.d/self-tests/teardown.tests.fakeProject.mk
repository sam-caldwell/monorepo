#
#test/Makefile/teardown.fakeProject.mk
# (c) 2023 Sam Caldwell.  See License.txt
#
test/Makefile/teardown.fakeProject:
	@$(PRINT_START)
	@$(RM_RF) cmd/fakeProject $(IGNORE_ERROR)
	@$(RM_RF)projects/fakeProject $(IGNORE_ERROR)
	@$(PRINT_DONE)