#
#test/Makefile/teardown.fakeProject.mk
# (c) 2023 Sam Caldwell.  See License.txt
#
test/Makefile/teardown.fakeProject:
	@@echo "\033[34m>starting $@\033[0m"
	@$(RM_RF) cmd/fakeProject $(IGNORE_ERROR)
	@$(RM_RF)projects/fakeProject $(IGNORE_ERROR)
	@echo "\033[32m>ok $@\033[0m"