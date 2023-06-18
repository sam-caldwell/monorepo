#
# self-tests/run.mk
# (c) 2023 Sam Caldwell.  See License.txt
#
.PHONY: test/self-tests/run
test/self-tests/run:
	@echo "$@ start"
	@make test/Makefile.d/self-tests/disable.tests
	@make test/Makefile.d/self-tests/enable.tests
	#ToDo: add more tests here.
	@echo "$@ done"
