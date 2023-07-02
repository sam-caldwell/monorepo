#
# Makefile.d/packer/build.packer.windows.10.mk
# (c) Sam Caldwell.  See LICENSE.txt
#

build/packer/windows/10:
	@echo "\033[34m>starting $@\033[0m"
	@$(PACKER_BIN) build -var-file $(PACKER_ROOT)/windows/10/main.pkrvars.hcl $(PACKER_ROOT)/windows/10/main.pkr.hcl;\
	@echo "\033[32m>ok $@\033[0m"