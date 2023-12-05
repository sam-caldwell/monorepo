# (c) 2022 Sam Caldwell.  All Rights Reserved.

TestReporter/install:
	@echo "## $@"
	@mkdir $(BIN_DIR) &> /dev/null || true
	@make $(shell dirname $@)/build
	@cp -vp $(ARTIFACT_PATH)/$$(dirname $@)/$$(dirname $@) $(BIN_DIR)/$$(dirname $@)
	@chmod +x $(BIN_DIR)/$$(dirname $@)
	@echo "$(BIN_DIR)/$$(dirname $@) ready"

