#
# list the projects that build artifacts.
# For build autoamation docs, see docs/builds/README.md
#
ifeq ($(OPSYS),windows)
	# Windows
	WINDOWS_LIST_PROJECTS = \
		echo ""; \
		echo "current binary projects (enabled):"; \
		for /d %%I in ($(BUILD_PROJECTS)) do echo " - %%I"; \
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


ifeq ($(OPSYS),windows)
list/build/projects:
	@$(PRINT_START)
	@$(WINDOWS_LIST_PROJECTS)
	@$(PRINT_DONE)
else
list/build/projects:
	@$(PRINT_START)
	@$(UNIX_LIST_PROJECTS)
	@$(PRINT_DONE)
endif

