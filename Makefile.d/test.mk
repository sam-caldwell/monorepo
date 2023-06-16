test: lint test/simple
	echo "$@ done"

test/simple:
	go test -v -count=1 ./...