#
#test/Makefile/teardown.fakeProject.mk
# (c) 2023 Sam Caldwell.  See License.txt
#
test/Makefile/teardown.fakeProject:
	@echo "\033[32m>starting $@\033[0m"
	@rm -rf cmd/fakeProject &> /dev/null || true
	@rm -rf projects/fakeProject &> /dev/null || true
	@echo "\033[32m>done:$@\033[0m"