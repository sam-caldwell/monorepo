PROJECTS := $(shell find cmd -mindepth 1 -maxdepth 1 -type d)
GOOS_LIST := windows linux darwin
GOARCH_LIST := amd64 arm64

build: $(patsubst cmd/%,build/%,$(PROJECTS))

.PHONY: build/$(PROJECTS)

build/%:
	@echo "Building $(subst cmd/,,$@)"
	@$(foreach GOOS,$(GOOS_LIST),\
		$(foreach GOARCH,$(GOARCH_LIST),\
			$(foreach COMMAND,$(shell find ./cmd/$(subst build/,,$@) -type d -mindepth 1 -maxdepth 1),\
				echo "Building $(GOOS)/$(GOARCH)/$(shell basename ${COMMAND})$(if $(findstring windows,$(GOOS)),.exe,)"; \
				GOOS=$(GOOS) GOARCH=$(GOARCH) go build -o ./build/$(GOOS)/$(GOARCH)/$(shell basename ${COMMAND})$(if $(findstring windows,$(GOOS)),.exe,) ./$(COMMAND)/main.go; \
			) \
		) \
	)
