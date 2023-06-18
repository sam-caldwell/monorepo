#
# list the projects that build artifacts.
# For build autoamation docs, see docs/builds/README.md
#
list/build/projects:
ifeq ($(OPSYS),windows)
	@echo off
	@echo "Windows fucking sucks.  If you are reading this, you've bought into the line that it's better.  It's not."
	@echo "I might fix this when I have time.  But...well...that would be enabling your dumb ass."
else
	@(\
		echo ""; \
		echo "current binary projects (enabled):"; \
		for PROJECT in $(BUILD_PROJECTS); do \
		  echo " - $${PROJECT}"; \
	    done; \
		echo ""; \
	)
endif