#
# Makefile.d/_config/gotmpdir.mk
# (c) Sam Caldwell.  See LICENSE.txt
#
# This is the temp directory which Go will use for it's operations (go run, go build, go test, etc).
# This is a workaround for a bug in golang which mostly affects Linux users who are attempting to
# build on Linux with NOEXEC turned on (as it should be) for /tmp directories.
#
# If this value is NOT set, golang will attempt to build/run binaries in /tmp.  If the Linux NOEXEC
# flag is turned on in /etc/fstab for /tmp, the resulting binary cannot be executed and the entire
# makefile system in this repository will fail.
#
# The following will create a new temp directory (build/tmp) and store the value in GOTMPDIR.
#
GOTMPDIR=$(shell mkdir -p build/tmp/ || true; echo build/tmp)
