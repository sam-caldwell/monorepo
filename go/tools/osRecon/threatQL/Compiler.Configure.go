package threatQL

import "github.com/sam-caldwell/monorepo/go/tools/osRecon/cli"

// Configure - Configure the object from environment variables and command-line arguments
func (compiler *Compiler) Configure() (err error) {
	if dir, err := cli.GetQueryCollectorDir(); err != nil {
		return err
	} else {
		compiler.queryDir = *dir
	}
	return err
}
