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
	@echo "\033[32m>>starting $@\033[0m"
	@(\
		export PROJECT="$(shell basename $@)"; \
		echo "disable project: $${PROJECT}";\
		export CMD_FLAG_FILE="./cmd/$${PROJECT}/test.disabled";\
		echo "CMD_FLAG_FILE:      $${CMD_FLAG_FILE}"; \
		\
		if [ -d $(shell dirname "$${CMD_FLAG_FILE}") ]; then \
			echo "disabling in cmd"; \
			if [ -f "$${CMD_FLAG_FILE}" ]; then \
				echo "project ($${PROJECT}) is already disabled in (cmd)"; \
			else \
				echo "$(shell whoami)@$(shell hostname) : $(shell date)" >> "$${CMD_FLAG_FILE}"; \
				echo "project ($${PROJECT}) disabled (cmd)";\
			fi; \
		fi; \
		\
		if [ -d $(shell dirname "$${PROJECTS_FLAG_FILE}") ]; then \
			echo "disabling in projects"; \
			if [ -f "$${PROJECTS_FLAG_FILE}" ]; then \
				echo "project ($${PROJECT}) is already disabled in (projects)"; \
			else \
				echo "$(shell whoami)@$(shell hostname) : $(shell date)" >> "$${PROJECTS_FLAG_FILE}"; \
				echo "project ($${PROJECT}) disabled (projects)";\
			fi; \
		fi; \
	)
