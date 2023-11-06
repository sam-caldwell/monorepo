package file

import (
	"fmt"
	"gopkg.in/yaml.v3"
	"os"
)

// IsValidYaml - Return boolean if YAML can be opened and unmarshalled without error
func IsValidYaml(filename string) error {
	yamlFile, err := os.ReadFile(filename)
	if err != nil {
		return fmt.Errorf("error reading YAML file '%s': %v", filename, err)
	}
	//
	var data any
	err = yaml.Unmarshal(yamlFile, &data)
	if err != nil {
		return fmt.Errorf("error unmarshaling YAML file '%s': %v", filename, err)
	}
	return nil
}
