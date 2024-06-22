# (c) 2022 Sam Caldwell.  All Rights Reserved.

application/Core/data:
	@echo "## $@"
	@mkdir -p "$(ARTIFACT_PATH)/" > /dev/null | true
	@cp -v $(PROJECT_ROOT)/$$(dirname $@)/tests/*.yaml "$(ARTIFACT_PATH)/$$(dirname $@)"
	@command -v yamllint &>/dev/null || {\
  		echo "prerequisite: yamllint not found"; \
  		exit 1; \
	}
	@yamllint "$(ARTIFACT_PATH)/$$(dirname $@)/test_config.yaml"
	@echo "test data is ready and in place"

application/Core/test:
	@echo "##"
	@echo "## $@"
	@echo "## Build logs directory: $(BUILD_LOG_DIRECTORY)"
	@echo "## Test name: $$(dirname $@)"
	@echo "##"
	@echo "Building tests...."
	@mkdir -p "$(ARTIFACT_PATH)/$$(dirname $@)" > /dev/null | true
	@mkdir -p "$(BUILD_LOG_DIRECTORY)/$$(dirname $@)" > /dev/null | true
	$(CPP_BUILD_COMMAND) -o $(ARTIFACT_PATH)/$$(dirname $@)/test $(PROJECT_ROOT)/$$(dirname $@)/test.cpp
	@make $$(dirname $@)/data
	@tree artifacts
	@echo "Running tests...."
	@( \
		cd "$(ARTIFACT_PATH)/$$(dirname $@)/" || exit 99; \
		./test "$(ARTIFACT_PATH)" "$$(dirname $@)" || exit $?; \
	) || exit $?

