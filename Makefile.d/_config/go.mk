#
# _config.go.mk
# (c) 2023 Sam Caldwell.  See LICENSE.txt
#
# This file sets GO_BINARY so we can accommodate
# the Redmond Mafia and every out-of-date technologist
# who things there was an ROI in me spending 2x as long
# to accommodate both their out-of-date shit-ware and
# the entirety of the POSIX-approximate world
#
ifeq ($(OPSYS),windows)
# since you use windows, here's a joke for you...
# You give Microsoft a bunch of money, and Microsoft
# gives you the finger...everytime.  You, therefore,
# are the joke.
GO_BINARY:="c:\Program Files\Go\bin\go.exe"
else
# We assume all other systems are rational posix systems
# who just call 'go' what it is...'go' because otherwise
# it is a group of sadistic, greedy, insecure people who
# probably work for Microsoft.
GO_BINARY:="go"
endif