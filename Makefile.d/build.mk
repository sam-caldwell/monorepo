#
# For build automation docs, see docs/builds/README.md
#
.PHONY: build
build:
	@echo "Building projects..."
	@(\
		for TARGET in $(BUILD_PROJECTS); do \
		    echo "  - Process TARGET: $${TARGET}"; \
			export PROGRAM="$$( basename $${TARGET})"; \
			echo "  - Process PROGRAM: $${PROGRAM}"; \
			export PROJECT="$$(basename $$( dirname $${TARGET} ))";\
			echo "  - Process PROJECT: $${PROJECT}"; \
			SOURCE="$${TARGET}/main.go"; \
			echo "  - Process SOURCE: $${SOURCE}"; \
			for GOOS in $(GOOS_LIST); do \
				for GOARCH in $(GOARCH_LIST); do \
					if [ "$${GOOS}_" == "windows_" ]; then \
					  export EXTENSION=".exe";\
					else \
					  export EXTENSION=""; \
					fi;\
				    export BINARY_ARTIFACT="$(BUILD_DIR)/$${GOOS}/$${GOARCH}/$${PROGRAM}$${EXTENSION}";\
					echo "  BUILD: $${BINARY_ARTIFACT} (source: $${SOURCE})"; \
					mkdir -p "$(BUILD_DIR)/$${GOOS}/$${GOARCH}"; \
					go build -o $${BINARY_ARTIFACT} "$${SOURCE}"; \
				done; \
			done; \
		done; \
	)
