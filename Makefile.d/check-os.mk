#
# 'make check-os'
# (c) 2023 Sam Caldwell.  MIT License.
#
# This Makefile is included in 'Makefile' in the root of the repo.
# It checks for an environment variable (OS) defined in Windows, and
# if that exists, it knows the operating system and we are done.
# But if the environment variable OS is not present, we know that we
# are probably on a POSIX system (Linux or MacOS--aka Darwin) because
# only the Redmond Mafia makes you jump through hoops to figure things
# out. (Translation: You can do better Microsoft...)
#
ifeq ($(OS),Windows_NT)
    DETECTED_OS := windows
else
    UNAME_S := $(shell uname -s)
    ifeq ($(UNAME_S),Linux)
        DETECTED_OS := linux
    else ifeq ($(UNAME_S),Darwin)
        DETECTED_OS := darwin
    else ifeq ($(UNAME_S),FreeBSD)
    	DETECTED_OS := freebsd
    else ifeq ($(UNAME_S),NetBSD)
    	DETECTED_OS := netbsd
	else ifeq ($(UNAME_S),OpenBSD)
		DETECTED_OS := openbsd
	else
		DETECED_OS := unknown
	endif
endif

check-os:
	@echo "$(DETECTED_OS)"
