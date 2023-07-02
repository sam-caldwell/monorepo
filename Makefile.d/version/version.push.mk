# make version/push
# (c) 2023 Sam Caldwell.  See LICENSE.txt
#
# This file defines 'make version/push' to push our new
# version to Github.
#
version/push:
	@echo "\033[34m>starting $@\033[0m"
	git push origin $(CURRENT_REPO_VERSION)
	@echo "\033[32m>ok $@\033[0m"
