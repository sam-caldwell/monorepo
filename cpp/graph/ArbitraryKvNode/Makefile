#
# (c) 2022 Sam Caldwell.  All Rights Reserved.
#
.PHONY: clean
clean:
	rm -rf build
	mkdir build

.PHONY: build
build:
	g++ -v -o build/ArbitraryKvNode main.cpp

.PHONY: run
run:
	./build/ArbitraryKvNode

.PHONY: build_test
build_test:
	g++ -v -o build/ArbitraryKvNode.test test.cpp

.PHONY: test
test: build_test
	./build/ArbitraryKvNode.test


build_dep_test:
	@echo "Build dependencies..."
	@(\
		cd ../ArbitraryKvNode;\
		make build_dep_test;\
	)
	@(\
		cd ../ArbitraryValue;\
		make test;\
	)
	@(\
		cd ../ArbitraryKeyValue;\
		make test;\
	)