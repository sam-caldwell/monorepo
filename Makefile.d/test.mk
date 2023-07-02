test: lint test/simple
	echo "$@ done"

test/simple: # test/Makefile.d/self-tests/run
	go test -failfast -v -count=1 -vet=all ./...