# (c) 2022 Sam Caldwell.  All Rights Reserved.
#
# This file will include the child projects (TcpServer, TcpClient, TcpCommon, etc)
#

include projects/networking/TcpClient/Makefile.d/*.mk
include projects/networking/TcpCommon/Makefile.d/*.mk
include projects/networking/TcpServer/Makefile.d/*.mk

networking/build:
	make networking/TcpCommon/build
	make networking/TcpClient/build
	make networking/TcpServer/build

networking/clean:
	make networking/TcpCommon/clean
	make networking/TcpClient/clean
	make networking/TcpServer/clean

networking/lint:
	make networking/TcpCommon/lint
	make networking/TcpClient/lint
	make networking/TcpServer/lint

networking/test:
	make networking/TcpCommon/test
	make networking/TcpClient/test
	make networking/TcpServer/test
