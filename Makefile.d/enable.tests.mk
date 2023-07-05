#
# enable.tests.mk
# (c) 2023 Sam Caldwell.  See License.txt
#
# This will enable 'make -s test' for a project in
# projects/ or cmd/.
#
# For build autoamation docs, see docs/builds/README.md
#
enable/tests/%:
	@echo "\033[32m>>starting $@\033[0m"
	@(\
		export PROJECT="$(shell basename $@)"; \
		echo "disable project: $${PROJECT}";\
		export CMD_FLAG_FILE="./cmd/$${PROJECT}/test.disabled";\
		export PROJECTS_FLAG_FILE="./projects/$${PROJECT}/test.disabled";\
		\
		if [ -d $(shell dirname "$${CMD_FLAG_FILE}") ]; then \
			if [ -f "$${CMD_FLAG_FILE}" ]; then \
				rm "$${CMD_FLAG_FILE}"; \
				echo "project enabled";\
			else \
				echo "project ($${PROJECT}) is already enabled in (cmd)"; \
			fi; \
		fi; \
		\
		if [ -d $(shell dirname "$${CMD_FLAG_FILE}") ]; then \
			if [ -f "$${PROJECTS_FLAG_FILE}" ]; then \
				rm "$${PROJECTS_FLAG_FILE}"; \
				echo "project enabled";\
			else \
				echo "project ($${PROJECT}) is already enabled in (projects)"; \
			fi; \
		fi; \
	)
