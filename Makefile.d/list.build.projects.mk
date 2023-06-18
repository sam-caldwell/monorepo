#
# list the projects that build artifacts.
# For build autoamation docs, see docs/builds/README.md
#
list/build/projects:
	@$(PRINT_START)
ifeq ($$(OPSYS),windows)
	@echo $(NOT_SUPPORTED)
	exit 1
else
	@echo ""; \
	echo "current binary projects (enabled):"; \
	for PROJECT in $(BUILD_PROJECTS); do \
		echo " - $$PROJECT"; \
	done; \
	echo "";
endif
	@$(PRINT_DONE)

