COLOR:=./bin/color


ifeq ($(OPSYS),windows)
    # Windows
    RM = del /Q /F
    MKDIR = mkdir
    ECHO = $(cmd /V:ON /C echo)

    IGNORE_ERROR=
    TERMINATE_ON_ERROR=
    PRINT_START=echo "start: $@ [OPSYS: $(OPSYS)]"

    ANSI_RED=^[[31m
    ANSI_GREEN=^[[32m
    ANSI_BLUE=^[[34m
    ANSI_RESET=^[[0m
else
    # Linux or other Unix-like systems
    RM = rm -f
    RM_RF = rm -rf
    MKDIR = mkdir -p
    ECHO=echo

	IGNORE_ERROR=&> /dev/null || true
	TERMINATE_ON_ERROR=|| { echo "FAILED:$@";exit 1; }
	PRINT_START=echo "$(ANSI_GREEN)start: $@ [OPSYS: $(OPSYS)]$(ANSI_RESET)"

	ANSI_RED="\\033[31m"
	ANSI_GREEN="\\033[32m"
	ANSI_BLUE="\\033[34m"
	ANSI_RESET="\\033[0m"
endif

