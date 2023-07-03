packer/init:
	@echo "\033[34m>starting $@\033[0m"
	@$(PACKER_BIN) init $(PACKER_ROOT)|| {\
		echo "\033[31m>failed $@\033[0m";\
		exit 1;\
	};\
	echo "\033[32m>ok $@\033[0m"