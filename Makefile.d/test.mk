test: lint test/simple
	echo "$@ done"

test/simple: test/Makefile.d/self-tests/run
	go test -v -count=1 ./...