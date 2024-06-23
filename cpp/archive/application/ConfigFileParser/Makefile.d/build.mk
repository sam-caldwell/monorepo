# (c) 2022 Sam Caldwell.  All Rights Reserved.

application/ConfigFileParser/build:
	@echo "## $@"
	@mkdir -p "$(ARTIFACT_PATH)/$$(dirname $@)" > /dev/null | true
	$(CPP_BUILD_COMMAND) -o $(ARTIFACT_PATH)/$$(dirname $@) $(PROJECT_ROOT)/$$(dirname $@)/main.cpp
	@echo "## $@ complete"
