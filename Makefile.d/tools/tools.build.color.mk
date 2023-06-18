tools/build/%:
	@echo "starting $@"
	@( \
		PROJECT=$(shell basename $(shell dirname $@));\
		PROGRAM=$(shell basename $@);\
		[ ! -f "./cmd/$${PROJECT}/$${PROGRAM}/main.go" ] && {\
			echo "file not found: ./cmd/$${PROJECT}/$${PROGRAM}/main.go";\
			exit 1;\
		};\
		go build -o "bin/$${PROGRAM}$${BUILD_EXTENSION}" "./cmd/$${PROJECT}/$${PROGRAM}/main.go";\
	)
	@echo "finished $@"
