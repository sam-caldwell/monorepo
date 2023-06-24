security/snyk/macos:
	brew tap snyk/tap
	brew install snyk

security/snyk:
	@snyk auth
	@snyk ignore --file-path='projects/argparse/Arguments.Copyright_test.go'
	@snyk code test

security: lint \
		  security/snyk
