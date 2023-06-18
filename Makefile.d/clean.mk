clean:
	@echo "clean start [OPSYS: $(OPSYS)]"
	@make test/Makefile/teardown.fakeProject
	@$(RM_RF) ./bin || { echo "FAILED:$@";exit 1; }
	@$(MKDIR) ./bin || { echo "FAILED:$@";exit 1; }
	@$(RM_RF) ./build || { echo "FAILED:$@";exit 1; }
	@$(MKDIR) ./build || { echo "FAILED:$@";exit 1; }
	@echo "clean...done"
