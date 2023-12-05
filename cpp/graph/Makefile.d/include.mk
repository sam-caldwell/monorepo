# (c) 2022 Sam Caldwell.  All Rights Reserved.

include projects/graph/Common/Makefile.d/*.mk
include projects/graph/ArbitraryKeyValue/Makefile.d/*.mk
include projects/graph/ArbitraryKvBtNode/Makefile.d/*.mk
include projects/graph/ArbitraryKvBtree/Makefile.d/*.mk
include projects/graph/ArbitraryKvList/Makefile.d/*.mk
include projects/graph/ArbitraryKvNode/Makefile.d/*.mk
include projects/graph/ArbitraryValue/Makefile.d/*.mk

graph/build:
	make graph/Common/build
	make graph/ArbitraryValue/build
	make graph/ArbitraryKvNode/build
	make graph/ArbitraryKeyValue/build
	make graph/ArbitraryKvList/build
	make graph/ArbitraryKvBtNode/build
	make graph/ArbitraryKvBtree/build

graph/clean:
	make graph/Common/clean
	make graph/ArbitraryValue/clean
	make graph/ArbitraryKvNode/clean
	make graph/ArbitraryKeyValue/clean
	make graph/ArbitraryKvList/clean
	make graph/ArbitraryKvBtNode/clean
	make graph/ArbitraryKvBtree/clean


graph/lint:
	make graph/Common/lint
	make graph/ArbitraryValue/lint
	make graph/ArbitraryKvNode/lint
	make graph/ArbitraryKeyValue/lint
	make graph/ArbitraryKvList/lint
	make graph/ArbitraryKvBtNode/lint
	make graph/ArbitraryKvBtree/lint

graph/test:
	make graph/Common/test
	make graph/ArbitraryValue/test
	make graph/ArbitraryKvNode/test
	make graph/ArbitraryKeyValue/test
	make graph/ArbitraryKvList/test
	make graph/ArbitraryKvBtNode/test
	make graph/ArbitraryKvBtree/test
