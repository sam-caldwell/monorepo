#
# _config.000.opsys.mk
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
    OPSYS := windows
else
    UNAME_S := $(shell uname -s)
    ifeq ($(UNAME_S),Linux)
        OPSYS := linux
    else ifeq ($(UNAME_S),Darwin)
        OPSYS := darwin
    else ifeq ($(UNAME_S),FreeBSD)
    	OPSYS := freebsd
    else ifeq ($(UNAME_S),NetBSD)
    	OPSYS := netbsd
	else ifeq ($(UNAME_S),OpenBSD)
		OPSYS := openbsd
	else
		OPSYS := unknown
	endif
endif
