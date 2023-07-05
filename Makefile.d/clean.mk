clean:
	$(PRINT_START)
	@make -s test/Makefile/teardown.fakeProject
	@$(RM_RF) $(REPO_ROOT)/bin $(TERMINATE_ON_ERROR)
	@$(MKDIR) $(REPO_ROOT)/bin $(TERMINATE_ON_ERROR)
	@$(RM_RF) $(REPO_ROOT)/build $(TERMINATE_ON_ERROR)
	@$(MKDIR) $(REPO_ROOT)/build $(TERMINATE_ON_ERROR)
	$(PRINT_DONE)
