# (c) 2022 Sam Caldwell.  All Rights Reserved.

TestReporter/test:
	@echo "## $@"
	@mkdir -p "$(ARTIFACT_PATH)/$$(dirname $@)" > /dev/null | true
	$(CPP_BUILD_COMMAND) -o $(ARTIFACT_PATH)/$$(dirname $@)/test $(PROJECT_ROOT)/$$(dirname $@)/test.cpp
	#@make $$(dirname $@)/data
	@tree artifacts
	@echo "Running tests...."
	@( \
		cd "$(ARTIFACT_PATH)/$$(dirname $@)/" || exit 99; \
		./test "$(ARTIFACT_PATH)" "$$(dirname $@)" || exit $?; \
	) || exit $?
