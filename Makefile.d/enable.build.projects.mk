#
# This will enable a cmd/<project> build
# For build autoamation docs, see docs/builds/README.md
#

enable/build/program/%:
	@(\
		export PROJECT="$(shell basename $@)"; \
		echo "disable project: $${PROJECT}";\
		if [[ ! -d "./cmd/$${PROJECT}" ]]; then \
			echo "invalid project ($${PROJECT})"; \
			exit 1; \
		else \
			if [[ -f "./cmd/$${PROJECT}/build.disabled" ]]; then \
				rm "./cmd/$${PROJECT}/build.disabled"; \
			else \
				echo "$${PROJECT} is already enabled"; \
			fi; \
		fi; \
	)
