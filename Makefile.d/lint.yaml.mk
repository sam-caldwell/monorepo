#
# To avoid trying to pull down yamllint across multiple platforms
# (ahem...looking at you Microsoft), I'm just gonna use this to do
# basic YAML validation.  Eventually I will create a schema validation
# feature.
#
# ToDo: Add Schema validator
#
YAML_VALIDATOR:=build/yamlValidator$(EXTENSION)

lint/yaml/setup:
	@go build -o $(YAML_VALIDATOR) ./cmd/validators/yamlValidator/main.go

lint/yaml: lint/yaml/setup
	./$(YAML_VALIDATOR) .
