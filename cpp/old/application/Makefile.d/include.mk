# (c) 2022 Sam Caldwell.  All Rights Reserved.
#
# This file will include the child projects (CommandLineParser, Application, Common, etc)
#

include projects/application/Common/Makefile.d/*.mk
include projects/application/CommandLineParser/Makefile.d/*.mk
include projects/application/ConfigFileParser/Makefile.d/*.mk
include projects/application/Configuration/Makefile.d/*.mk
include projects/application/Core/Makefile.d/*.mk
include projects/application/EnvironmentVariableParser/Makefile.d/*.mk

application/build:
	make application/Common/build
	make application/Configuration/build
	make application/CommandLineParser/build
	make application/ConfigFileParser/build
	make application/EnvironmentVariableParser/build
	make application/Core/build

application/clean:
	make application/Common/clean
	make application/Configuration/clean
	make application/CommandLineParser/clean
	make application/ConfigFileParser/clean
	make application/EnvironmentVariableParser/clean
	make application/Core/clean


application/lint:
	make application/Common/lint
	make application/Configuration/lint
	make application/CommandLineParser/lint
	make application/ConfigFileParser/lint
	make application/EnvironmentVariableParser/lint
	make application/Core/lint

application/test:
	make application/Common/test
	make application/Configuration/test
	make application/CommandLineParser/test
	make application/ConfigFileParser/test
	make application/EnvironmentVariableParser/test
	make application/Core/test