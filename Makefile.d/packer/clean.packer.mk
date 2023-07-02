clean/packer:
	$(PRINT_START)
	@rm -rf $(PACKER_BIN) &>/dev/null || true
	$(PRINT_DONE)