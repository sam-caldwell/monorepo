#
# (c) 2022 Sam Caldwell.  All Rights Reserved.
#
.PHONY: clean
clean:
	rm -rf build
	mkdir build

.PHONY: build
build:
	g++ -v -o build/SmartPointer main.cpp

.PHONY: run
run:
	./build/SmartPointer

.PHONY: build_test
build_test: build_dep_test
	g++ -v -o build/SmartPointer.test test.cpp


.PHONY: test
test: build_test
	./build/SmartPointer.test

build_dep_test:
	@echo "No dependencies"
	@exit 0