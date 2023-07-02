clean:
	@echo "\033[34m>starting $@\033[0m"
	@make test/Makefile/teardown.fakeProject
	@$(RM_RF) $(REPO_ROOT)/bin $(TERMINATE_ON_ERROR)
	@$(MKDIR) $(REPO_ROOT)/bin $(TERMINATE_ON_ERROR)
	@$(RM_RF) $(REPO_ROOT)/build $(TERMINATE_ON_ERROR)
	@$(MKDIR) $(REPO_ROOT)/build $(TERMINATE_ON_ERROR)
	@make clean/packer
	echo "\033[32m>ok $@\033[0m"
