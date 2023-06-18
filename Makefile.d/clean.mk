clean:
	@$(PRINT_START)
	@make test/Makefile/teardown.fakeProject
	@$(RM_RF) ./bin $(TERMINATE_ON_ERROR)
	@$(MKDIR) ./bin $(TERMINATE_ON_ERROR)
	@$(RM_RF) ./build $(TERMINATE_ON_ERROR)
	@$(MKDIR) ./build $(TERMINATE_ON_ERROR)
	@echo "done: $@"
