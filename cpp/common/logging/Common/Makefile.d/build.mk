# (c) 2022 Sam Caldwell.  All Rights Reserved.

logging/Common/build:
	@echo "## $@"
	@mkdir -p "$(ARTIFACT_PATH)/$$(dirname $@)" > /dev/null | true
	$(CPP_BUILD_COMMAND) -o $(ARTIFACT_PATH)/$$(dirname $@)/$$(dirname $@) $(PROJECT_ROOT)/$$(dirname $@)/main.cpp
	@echo "## $@ complete"
