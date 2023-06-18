#
# For build automation docs, see docs/builds/README.md
#
.PHONY: build
build:
	@echo "\033[32m>starting $@\033[0m"
	@(\
		for TARGET in $(BUILD_PROJECTS); do \
		    echo "\033[32m  - Process TARGET: $${TARGET}"; \
			export PROGRAM="$$( basename $${TARGET})"; \
			echo "\033[32m  - Process PROGRAM: $${PROGRAM}"; \
			export PROJECT="$$(basename $$( dirname $${TARGET} ))";\
			echo "\033[32m  - Process PROJECT: $${PROJECT}"; \
			SOURCE="$${TARGET}/main.go"; \
			echo "\033[32m  - Process SOURCE: $${SOURCE}"; \
			for GOOS in $(GOOS_LIST); do \
				for GOARCH in $(GOARCH_LIST); do \
					if [ "$${GOOS}_" == "windows_" ]; then \
					  export BUILD_EXTENSION=".exe";\
					else \
					  export BUILD_EXTENSION=""; \
					fi;\
				    export BINARY_ARTIFACT="$(BUILD_DIR)/$${GOOS}/$${GOARCH}/$${PROGRAM}$${BUILD_EXTENSION}";\
					echo "\033[32m  BUILD: $${BINARY_ARTIFACT} (source: $${SOURCE})"; \
					mkdir -p "$(BUILD_DIR)/$${GOOS}/$${GOARCH}"; \
					go build -o $${BINARY_ARTIFACT} "$${SOURCE}"; \
				done; \
			done; \
		done; \
		echo "\033[32m>clean...done\033[0m"
	)
