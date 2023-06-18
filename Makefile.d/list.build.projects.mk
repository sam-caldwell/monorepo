#
# list the projects that build artifacts.
# For build autoamation docs, see docs/builds/README.md
#
ifeq ($(OPSYS),windows)
	# Windows
	WINDOWS_LIST_PROJECTS = \
		echo ""; \
		echo "current binary projects (enabled):"; \
		for %%I in ($(BUILD_PROJECTS)) do echo " - %%~I"; \
		echo "";
else
	# Unix-like systems
	UNIX_LIST_PROJECTS = \
		echo ""; \
		echo "current binary projects (enabled):"; \
		for PROJECT in $(BUILD_PROJECTS); do \
			echo " - $$PROJECT"; \
		done; \
		echo "";
endif

list/build/projects:
	@$(PRINT_START)
ifeq ($(OPSYS),windows)
	@$(WINDOWS_LIST_PROJECTS)
else
	@$(UNIX_LIST_PROJECTS)
endif
	@$(PRINT_DONE)
