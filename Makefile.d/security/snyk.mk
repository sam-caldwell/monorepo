#
# Note: To ignore a file, use this pattern
# 		snyk ignore --file-path='projects/argparse/Arguments.Copyright_test.go'
#
security/snyk: security/snyk/install
	@echo "\033[34m>starting $@\033[0m"
	@snyk code test
	@echo "\033[32m>ok $@\033[0m"
