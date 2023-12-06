package monorepo

import (
	"github.com/sam-caldwell/monorepo/go/ansi"
	"gopkg.in/yaml.v3"
	"os"
)

func (m *Manifest) Load(fileName string) (err error) {
	m.FileName = fileName

	ansi.Cyan().Printf("Load Manifest: %s\n", fileName).Reset()

	yamlFile, err := os.ReadFile(fileName)
	if err != nil {
		return err
	}

	err = yaml.Unmarshal(yamlFile, &m.config)
	if err != nil {
		return err
	}

	return nil
}
