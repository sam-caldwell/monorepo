#
# disable.tests.mk
# (c) 2023 Sam Caldwell.  See License.txt
#
# This will disable 'make test' for a project in
# projects/ or cmd/.
#
# For build autoamation docs, see docs/builds/README.md
#
disable/tests/%:
	@(\
		export PROJECT="$(shell basename $@)"; \
		echo "disable project: $${PROJECT}";\
		export CMD_FLAG_FILE="./cmd/$${PROJECT}/test.disabled"
		export PROJECTS_FLAG_FILE="./projects/$${PROJECT}/test.disabled"
		if [[ -d "$${CMD_FLAG_FILE}" ]]; then \
			if [[ -f "$${CMD_FLAG_FILE}" ]]; then \
				echo "project ($${PROJECT}) is already disabled in (cmd)"; \
			else \
				echo "$(shell whoami)@$(shell hostname) : $(shell date)" >> "$${CMD_FLAG_FILE}"; \
			fi; \
		fi; \
		if [[ -d "$${PROJECTS_FLAG_FILE}" ]]; then \
			if [[ -f "$${CMD_FLAG_FILE}" ]]; then \
				echo "project ($${PROJECT}) is already disabled in (projects)"; \
			else \
				echo "$(shell whoami)@$(shell hostname) : $(shell date)" >> "$${CMD_FLAG_FILE}"; \
			fi; \
		fi; \
	)
