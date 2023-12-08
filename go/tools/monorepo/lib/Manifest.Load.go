package monorepo

import (
	"github.com/sam-caldwell/monorepo/go/ansi"
	"gopkg.in/yaml.v3"
	"os"
)

// Load - Load the YAML Manifest file for the monorepo command
func (m *Manifest) Load(FileName *string) (err error) {

	rawContent, err := os.ReadFile(*FileName)
	if err != nil {
		ansi.Red().Printf("Error Loading Manifest: %s\n%v\n", *FileName, err).Reset()
		return err
	}

	err = yaml.Unmarshal(rawContent, &m.config)
	if err != nil {
		ansi.Red().Printf("Error Parsing Manifest: %s\n%v\n", *FileName, err).Reset()
		return err
	}

	return nil
}
