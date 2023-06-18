#
# list the projects that build artifacts.
# For build autoamation docs, see docs/builds/README.md
#
list/build/projects:
ifeq ($(OPSYS),windows)
	@$(NOT_SUPPORTED)
	exit 1
endif
	@$(PRINT_START)
	@echo ""; \
	echo "current binary projects (enabled):"; \
	for PROJECT in $(BUILD_PROJECTS); do \
		echo " - $$PROJECT"; \
	done; \
	echo "";
	@$(PRINT_DONE)

