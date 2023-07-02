# make version/current
# (c) 2023 Sam Caldwell.  See LICENSE.txt
#
# This file defines 'make version/current' to display the current
# repo version from the perspective of our Makefile system
#

version/current:
	@echo "current version: $(CURRENT_REPO_VERSION)"