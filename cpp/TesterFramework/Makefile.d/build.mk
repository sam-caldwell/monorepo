# (c) 2022 Sam Caldwell.  All Rights Reserved.

Tester/build:
	@echo "## $@ has no need to test."
	@mkdir -p "$(ARTIFACT_PATH)/$$(dirname $@)" > /dev/null | true
	@exit 0
