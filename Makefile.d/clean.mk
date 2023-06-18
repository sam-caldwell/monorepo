clean:
	@make test/Makefile/teardown.fakeProject
	@rm -rf ./build
	@mkdir -p ./build
