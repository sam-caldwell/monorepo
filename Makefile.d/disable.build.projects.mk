#
# This will disable a cmd/<project> build
# For build autoamation docs, see docs/builds/README.md
#
disable/build/program/%:
	@(\
		export PROJECT="$(shell basename $@)"; \
		echo "disable project: $${PROJECT}";\
		if [[ ! -d "./cmd/$${PROJECT}" ]]; then \
			echo "invalid project ($${PROJECT})"; \
			exit 1; \
		else \
			if [[ -f "./cmd/$${PROJECT}/build.disabled" ]]; then \
				echo "project ($${PROJECT}) is already disabled"; \
			else \
				echo "$(shell whoami)@$(shell hostname) : $(shell date)" >> "./cmd/$${PROJECT}/build.disabled"; \
			fi; \
		fi; \
	)
