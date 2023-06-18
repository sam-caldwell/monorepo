#
# list the projects that build artifacts.
# For build autoamation docs, see docs/builds/README.md
#
list/build/projects:
	@$(NOT_SUPPORTED)
	@$(PRINT_START)
	@echo ""; \
	echo "current binary projects (enabled):"; \
	for PROJECT in $(BUILD_PROJECTS); do \
		echo " - $$PROJECT"; \
	done; \
	echo "";
	@$(PRINT_DONE)

