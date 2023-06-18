ifeq ($(OPSYS),windows)
BUILD_PROJECTS := $(shell for /r "cmd" %G in (.) do @(where /q "%%~G\build.disabled" || echo %%~dpG))
else
BUILD_PROJECTS := $(shell find cmd -mindepth 2 -maxdepth 3 -type d -execdir test ! -f "build.disabled" \; -print)
endif
