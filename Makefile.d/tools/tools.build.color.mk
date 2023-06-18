tools/build/%:
	@echo "starting $@"
	@( \
		TOOL_PROJECT=$(shell basename $(shell dirname $@));\
		TOOL_PROGRAM=$(shell basename $@);\
		[ ! -f "./cmd/$${TOOL_PROJECT}/$${TOOL_PROGRAM}/main.go" ] && {\
			echo "file not found: ./cmd/$${TOOL_PROJECT}/$${TOOL_PROGRAM}/main.go";\
			exit 1;\
		};\
		go build -o "bin/$${TOOL_PROGRAM}$${BUILD_EXTENSION}" "./cmd/$${TOOL_PROJECT}/$${TOOL_PROGRAM}/main.go";\
	)
	@echo "finished $@"
