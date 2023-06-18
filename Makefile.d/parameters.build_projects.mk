ifeq ($(OPSYS),windows)
BUILD_PROJECTS := $(shell for /r "cmd" %G in (.) do @(dir /b /ad "%%~G" | findstr /r /c:"\\[^\\]*\\[^\\]*$$" >nul || echo %%~dpG))
else
BUILD_PROJECTS := $(shell find cmd -mindepth 2 -maxdepth 3 -type d -execdir test ! -f "build.disabled" \; -print)
endif
