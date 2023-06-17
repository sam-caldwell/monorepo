#
# list the projects that build artifacts.
# For build autoamation docs, see docs/builds/README.md
#
list/build/projects:
	@(\
		echo ""; \
		echo "current binary projects (enabled):"; \
		for PROJECT in $(PROJECTS); do \
		  echo " - $${PROJECT}"; \
	    done; \
		echo ""; \
	)
