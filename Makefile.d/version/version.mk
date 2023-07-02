# make version
# (c) 2023 Sam Caldwell.  See LICENSE.txt
#
# This file defines 'make version' to bump the version tag on the repo
# as well as the SVG image file and push the result to the current branch
# as a tagged new release candidate.
#

version: version/bump \
		 version/badge \
		 version/commit \
		 version/push
	@echo "\033[32m>ok $@\033[0m"
