# (c) 2022 Sam Caldwell.  All Rights Reserved.

common/exceptions/test:
	@echo "##"
	@echo "## $@"
	@echo "## Build logs directory: $(BUILD_LOG_DIRECTORY)"
	@echo "## Test name: $$(dirname $@)"
	@echo "##"
	@echo "Building tests...."
	@mkdir -p "$(ARTIFACT_PATH)/$$(dirname $@)" > /dev/null | true
	#@mkdir -p "$(BUILD_LOG_DIRECTORY)/$$(dirname $@)" > /dev/null | true
	$(CPP_BUILD_COMMAND) -o $(ARTIFACT_PATH)/$$(dirname $@)/test $(PROJECT_ROOT)/$$(dirname $@)/test.cpp
	@tree artifacts
	@echo "Running tests...."
	@( \
		cd "$(ARTIFACT_PATH)/$$(dirname $@)/" || exit 99; \
		./test "$(ARTIFACT_PATH)" "$$(dirname $@)" || exit $?; \
	) || exit $?

