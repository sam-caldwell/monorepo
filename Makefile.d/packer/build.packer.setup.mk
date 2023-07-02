#
# Makefile.d/packer/build.packer.setup.mk
# (c) Sam Caldwell.  See LICENSE.txt
#
build/packer/setup:
	@echo "\033[34m>starting $@\033[0m"
	@command -v $(PACKER_BIN) &> /dev/null || {\
  		echo "\033[34m>packer not found.  installing\033[0m";\
  		mkdir -p bin/ || true;\
		curl --fail \
			 -s \
			 -L https://releases.hashicorp.com/packer/$(PACKER_VERSION)/packer_$(PACKER_VERSION)_$(OPSYS)_$(CPU_ARCH).zip \
			 -o build/packer_$(PACKER_VERSION)_$(OPSYS)_$(CPU_ARCH).zip; \
		unzip build/packer_$(PACKER_VERSION)_$(OPSYS)_$(CPU_ARCH).zip \
			  -d $(BIN_DIR)/; \
		rm build/packer_$(PACKER_VERSION)_$(OPSYS)_$(CPU_ARCH).zip; \
	} || {\
		echo "\033[31m>failed $@\033[0m";\
		exit 1;\
	};\
	echo "\033[32m>ok $@\033[0m"