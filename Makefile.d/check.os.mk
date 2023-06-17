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
	#
	# Special code for the special people who are stuck in the '90s
	#
    DETECTED_OS := windows
    GO_BINARY:="c:\Program Files\Go\bin\go.exe"

    # mac/linux package managers
    HAS_BREW:=no

    # linux: debian/ubuntu
    HAS_APT:=no
    HAS_DPKG:=no

    # linux: redhat/centos/fedora
    HAS_YUM:=no

    # linux: redhat/centos/fedora/sles
    HAS_RPM:=no
else
    UNAME_S := $(shell uname -s)
    GO_BINARY:="go"
	#
	# PACKAGE MANAGERS
	#
	# mac/linux package managers
	HAS_BREW=$(shell command -v brew &> /dev/null && echo yes)
	# linux: debian/ubuntu
	HAS_APT=$(shell command -v apt-get &> /dev/null && echo yes)
	HAS_DPKG=$(shell command -v dpkg &> /dev/null && echo yes)

	# linux: redhat/centos/fedora
	HAS_YUM=$(shell command -v yum &> /dev/null && echo yes)

	# linux: redhat/centos/fedora/sles
	HAS_RPM=$(shell command -v rpm &> /dev/null && echo yes)
	#
	# POSIX OPERATING SYSTEMS
	#
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

#
# PACKAGE MANAGERS
#
# windows package manager options
#HAS_WINGET:=$(shell "c:\Program Files\Go\bin\go.exe" run "Makefile.d/tools/has_winget.go")



HAS_WINGET:=$(shell ${GO_BINARY} run Makefile.d/tools/has_executable.go winget)
has_winget:
	@echo "has_winget $(HAS_WINGET)"

HAS_BREW:=$(shell ${GO_BINARY} run Makefile.d/tools/has_executable.go brew)
has_brew:
	@echo "has_brew $(HAS_BREW)"
