#
# Makefile
# (c) 2023 Sam Caldwell.  See License.txt
#
# This is the root of our Makefile system.  It includes our Makefiles from Makefiles.d/
# and our Makefile tests in Makefiles.d/self-tests

include Makefile.d/*.mk
#
# To run the self-tests, we execute
# make test/self-tests/run
#
include Makefile.d/self-tests/*.mk