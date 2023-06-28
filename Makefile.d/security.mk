security/snyk/install/brew:
	@echo "\033[34m>starting $@\033[0m"
	@brew tap snyk/tap || { \
	  echo "\033[31m>failed $@\033[0m";\
	  exit 1;\
	};\
	export HOMEBREW_NO_INSTALL_CLEANUP=1;\
	brew reinstall snyk || { \
	  echo "\033[31m>failed $@\033[0m";\
	  exit 1;\
	};\
	brew cleanup snyk || { \
	  echo "\033[31m>failed $@\033[0m";\
	  exit 1;\
	};\
	echo "\033[32m>ok $@\033[0m"

security/snyk/install/winget:
	@echo "\033[34m>$@ not implemented yet\033[0m"
	@exit 1

security/snyk/install/choco:
	@echo "\033[34m>$@ not implemented yet\033[0m"
	@exit 1

security/snyk/install/apt:
	@echo "\033[34m>$@ not implemented yet\033[0m"
	@exit 1

security/snyk/install/yum:
	@echo "\033[34m>$@ not implemented yet\033[0m"
	@exit 1

security/snyk/install:
	@echo "\033[34m>starting $@\033[0m"
	@command -v snyk &>/dev/null || make security/snyk/install/$(PACKAGE_MANAGER)
	@echo "\033[32m>ok $@\033[0m"

security/snyk/auth:
	@echo "\033[34m>starting $@\033[0m"
	@snyk auth
	@echo "\033[32m>ok $@\033[0m"

security/snyk: security/snyk/install
	@echo "\033[34m>starting $@\033[0m"
	#snyk ignore --file-path='projects/argparse/Arguments.Copyright_test.go'
	@snyk code test
	@echo "\033[32m>ok $@\033[0m"

security: lint \
		  security/snyk
	@echo "\033[32m>ok $@\033[0m"