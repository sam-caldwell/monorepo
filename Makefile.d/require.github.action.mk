# make require-github-action
# (c) 2023 Sam Caldwell.  See LICENSE.txt
#
# This file defines 'make version' to bump the version tag on the repo
# as well as the SVG image file and push the result to the current branch
# as a tagged new release candidate.
#
require-github-action:
	@if [ -z "$$GITHUB_ACTIONS" ]; then \
		echo "This command should not be run outside of GitHub Actions."; \
		exit 99; \
	fi
