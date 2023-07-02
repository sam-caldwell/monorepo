
build/packer/linux/ubuntu:
	@$(PACKER_BIN) build -var-file windows.10.pkrvars.hcl windows.10.pkr.hcl
	@echo "\033[32m>ok $@\033[0m"

