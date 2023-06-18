#
#test/Makefile/teardown.fakeProject.mk
# (c) 2023 Sam Caldwell.  See License.txt
#
test/Makefile/teardown.fakeProject:
	@echo "starting $@"
	@$(RM_RF) cmd/fakeProject &> /dev/null || true
	@$(RM_RF)projects/fakeProject &> /dev/null || true
	@echo "done: $@"