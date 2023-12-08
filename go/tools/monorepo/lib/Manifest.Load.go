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

	err = yaml.Unmarshal(rawContent, &m.config)
	if err != nil {
		ansi.Red().Printf("Error Parsing Manifest: %s\n%v\n", *FileName, err).Reset()
		return err
	}
	if m.FileName != *FileName {
		ansi.Red().Printf("Filename did not load properly (%s)", *FileName).
			LF().Reset().Fatal(exit.GeneralError)
	}
	if m.config.Header.Author == "" {
		ansi.Red().Printf("Author did not load properly (%s)", *FileName).
			LF().Reset().Fatal(exit.GeneralError)
	}
	if m.config.Header.Email == "" {
		ansi.Red().Printf("Email did not load properly (%s)", *FileName).
			LF().Reset().Fatal(exit.GeneralError)
	}
	if m.config.Header.Description == "" {
		ansi.Red().Printf("Description did not load properly (%s)", *FileName).
			LF().Reset().Fatal(exit.GeneralError)
	}
	if len(m.config.Header.Arch) == 0 {
		ansi.Red().Printf("Arch did not load properly (%s)", *FileName).
			LF().Reset().Fatal(exit.GeneralError)
	}
	if len(m.config.Header.Opsys) == 0 {
		ansi.Red().Printf("Opsys did not load properly (%s)", *FileName).
			LF().Reset().Fatal(exit.GeneralError)
	}
	return nil
}
