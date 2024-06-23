# (c) 2022 Sam Caldwell.  All Rights Reserved.

graph/ArbitraryKvBtree/clean:
	@echo "clean..."
	@rm -rf $(ARTIFACT_PATH)/$$(dirname $@)
	@mkdir -p $(ARTIFACT_PATH)/$$(dirname $@)
	@echo "## $@ complete"
