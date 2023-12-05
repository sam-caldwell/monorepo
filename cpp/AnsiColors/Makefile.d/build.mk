# (c) 2022 Sam Caldwell.  All Rights Reserved.

AnsiColors/build:
	@echo "## $@ has no need to build."
	@mkdir -p "$(ARTIFACT_PATH)/$$(dirname $@)" > /dev/null | true
	@exit 0
