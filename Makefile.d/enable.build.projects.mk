#
# enable.build.projects.mk
# (c) 2023 Sam Caldwell.  See License.txt
#
# This will enable a cmd/<project> build
# For build autoamation docs, see docs/builds/README.md
#
enable/build/%:
	@echo "\033[32m>>starting $@\033[0m"
	@(\
		export PROJECT="$(shell basename $@)"; \
		export CMD_FLAG_FILE="./cmd/$${PROJECT}/build.disabled";\
		@echo "\033[34m enable project: '$${PROJECT}' $@\033[0m";\
		if [[ ! -d "$(shell dirname \\"$${CMD_FLAG_FILE}\\")" ]]; then \
			echo "\033[31m>invalid project ($${PROJECT}) $@\033[0m"; \
			exit 1; \
		else \
			if [[ -f "./cmd/$${PROJECT}/build.disabled" ]]; then \
				rm "./cmd/$${PROJECT}/build.disabled"; \
			else \
				echo "\033[31m>project $${PROJECT} is already enabled $@\033[0m"; \
			fi; \
		fi; \
	)
