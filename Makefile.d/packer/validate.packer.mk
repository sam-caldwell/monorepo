validate/packer: packer/init
	@echo "\033[34m>starting $@\033[0m"
	@while read -r opsys os_version os_family disabled; do \
		if [ "$${disabled}_" == "_" ]; then \
			$(PACKER_BIN) validate -var opsys=$${opsys} \
								   -var os_version=$${os_version} \
								   -var os_family="$${os_family}" systems/packer/ || {\
              		echo "\033[31m>>$${opsys} $${os_version}: validation failed\033[0m";\
              		exit 1;\
              	};\
		else\
			echo "\033[33m>>$${opsys} $${os_version}: disabled\033[0m";\
		fi; \
	done < systems/packer/supported_systems.txt
	@echo "\033[32m>ok $@\033[0m"
