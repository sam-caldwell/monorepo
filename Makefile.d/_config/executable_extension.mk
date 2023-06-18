#
# Makefile.d/_config/executable_extension.mk
# (c) 2023 Sam Caldwell.  See LICENSE.txt
#
# This parameter determines what (if any) binary
# extension we will append to the end of executables
#
ifeq ($(OS),Windows_NT)
    EXECUTABLE_EXTENSION := ".exe"
else
	EXECUTABLE_EXTENSION := ""
endif
