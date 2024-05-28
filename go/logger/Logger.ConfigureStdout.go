package logger

import (
	configuration "github.com/sam-caldwell/monorepo/go/configuration/Map"
	"os"
)

// ConfigureStdout - Configure logging to stdout
//
//	(c) 2023 Sam Caldwell.  MIT License
func (log *Logger) ConfigureStdout(cfg *configuration.Map[string, string]) {
	if cfg != nil {
		//Do something to use the configuration
	}
	log.target = os.Stdout
}
