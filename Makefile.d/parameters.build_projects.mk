ifeq ($(OPSYS),windows)
BUILD_PROJECTS := $(shell powershell -NoProfile -Command "Get-ChildItem -Path 'cmd' -Recurse -Directory | ForEach-Object { if (-not (Test-Path -Path (Join-Path $_.FullName 'build.disabled'))) { $_.FullName } }")
else
BUILD_PROJECTS := $(shell find cmd -mindepth 2 -maxdepth 3 -type d -execdir test ! -f "build.disabled" \; -print)
endif
