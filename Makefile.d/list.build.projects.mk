#
# list the projects that build artifacts.
# For build autoamation docs, see docs/builds/README.md
#
ifeq ($(OPSYS),windows)
list/build/projects:
	@$(WINDOWS_NOT_SUPPORTED)
else
list/build/projects:
	@$(PRINT_START)
	echo ""; \
	echo "current binary projects (enabled):"; \
	for PROJECT in $(BUILD_PROJECTS); do \
		echo " - $$PROJECT"; \
	done; \
	echo "";
	@$(PRINT_DONE)
endif

