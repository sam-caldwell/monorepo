#
# list the projects that build artifacts.
# For build autoamation docs, see docs/builds/README.md
#
list/build/projects:
	@$(PRINT_START)
ifeq ($(OPSYS),windows)
	@powershell -NoProfile -Command "Write-Host 'current binary projects (enabled):'; $(foreach PROJECT,$(BUILD_PROJECTS), Write-Host ' - $(PROJECT)')"
else
	@( \
		echo ""; \
		echo "current binary projects (enabled):"; \
		for PROJECT in $(BUILD_PROJECTS); do \
			echo " - $$PROJECT"; \
		done; \
		echo ""; \
	)
endif
	@$(PRINT_DONE)