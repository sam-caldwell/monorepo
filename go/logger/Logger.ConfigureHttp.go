package logger

import (
	configuration "github.com/sam-caldwell/monorepo/go/configuration/Map"
	"github.com/sam-caldwell/monorepo/go/exit/errors"
)

// ConfigureHttp - Configure a http network connection
//
//	(c) 2023 Sam Caldwell.  MIT License
func (log *Logger) ConfigureHttp(cfg *configuration.Map[string, string]) {
	//ToDo: implement syslog network io.writer
	log.useColor = false
	panic(errors.NotImplemented)
}
