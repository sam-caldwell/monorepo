clean:
	@echo "$(ANSI_GREEN)>starting$(ANSI_RESET)"
	@make test/Makefile/teardown.fakeProject &> /dev/null
	@$(RM_RF) ./bin || { echo -e "$(ANSI_RED)>>FAILED:$(ANSI_RESET)";exit 1; }
	@$(MKDIR) ./bin || { echo -e "$(ANSI_RED)>>FAILED:$(ANSI_RESET)";exit 1; }
	@$(RM_RF) ./build || { echo -e "$(ANSI_RED)>>FAILED:$(ANSI_RESET)";exit 1; }
	@$(MKDIR) ./build || { echo -e "$(ANSI_RED)>>FAILED:$(ANSI_RESET)";exit 1; }
	@echo "$(ANSI_GREEN)>clean...done$(ANSI_RESET)"
