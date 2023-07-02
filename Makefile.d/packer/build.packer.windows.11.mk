#
# Makefile.d/packer/build.packer.windows.11.mk
# (c) Sam Caldwell.  See LICENSE.txt
#
build/packer/windows/11:
	@$(PACKER_BIN) build -var-file windows.11.pkrvars.hcl windows.11.pkr.hcl
	@echo "\033[32m>ok $@\033[0m"