clean/packer:
	@echo "\033[34m>starting $@\033[0m"
	@rm -rf $(PACKER_BIN) &>/dev/null || true
	echo "\033[32m>ok $@\033[0m"
