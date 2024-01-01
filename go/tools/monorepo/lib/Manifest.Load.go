package monorepo

import (
	"github.com/sam-caldwell/monorepo/go/ansi"
	"github.com/sam-caldwell/monorepo/go/exit"
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
	if err = yaml.Unmarshal(rawContent, &m.config); err != nil {
		ansi.Red().Printf("Error Parsing Manifest: %s\n%v\n", *FileName, err).Reset()
		return err
	}
	if m.FileName != *FileName {
		ansi.Red().Printf("Filename did not load properly (%s)", *FileName).
			LF().Reset().Fatal(exit.GeneralError)
	}
	if err = m.Verify(); err != nil {
		ansi.Red().Printf("%v (%s)", err, *FileName).
			LF().Reset().Fatal(exit.GeneralError)
	}
	return err
}
