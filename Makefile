#
# Makefile
# (c) 2023 Sam Caldwell.  See License.txt
#
# This is the root of our Makefile system.  It includes our Makefiles from Makefiles.d/
# and our Makefile tests in Makefiles.d/self-tests

#
# _config contains variable settings,
# some of which are opsys-specific
#
include Makefile.d/*.mk

#
# To run linters, we execute
# make lint
#
include Makefile.d/lint/*.mk

#
# Checks are make targets that tell
# us about the initialized make environment
# we are running in.
#
include Makefile.d/check/*.mk

#
# To build or run our tools, we have
# these makefiles
#
include Makefile.d/tools/*.mk