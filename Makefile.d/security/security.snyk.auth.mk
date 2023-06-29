security/snyk/auth:
	@echo "\033[34m>starting $@\033[0m"
	@snyk auth
	@echo "\033[32m>ok $@\033[0m"