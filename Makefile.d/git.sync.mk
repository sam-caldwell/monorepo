git/sync:
	git fetch origin $(GIT_CURRENT_BRANCH)
	git rebase origin/$(GIT_CURRENT_BRANCH)
	git push origin $(GIT_CURRENT_BRANCH)