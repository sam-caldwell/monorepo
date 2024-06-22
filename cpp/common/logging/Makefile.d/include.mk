# (c) 2022 Sam Caldwell.  All Rights Reserved.
#
# This file will include the child projects (LogServer, Logger, Common, etc)
#

include projects/logging/Logger/Makefile.d/*.mk
include projects/logging/Common/Makefile.d/*.mk
include projects/logging/LogServer/Makefile.d/*.mk

logging/build:
	make logging/Common/build
	make logging/Logger/build
	make logging/LogServer/build

logging/clean:
	make logging/Common/clean
	make logging/Logger/clean
	make logging/LogServer/clean

logging/lint:
	make logging/Common/lint
	make logging/Logger/lint
	make logging/LogServer/lint

logging/test:
	make logging/Common/test
	make logging/Logger/test
	make logging/LogServer/test
