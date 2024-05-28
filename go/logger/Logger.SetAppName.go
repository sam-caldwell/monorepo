package logger

import "github.com/sam-caldwell/monorepo/go/exit/errors"

// SetAppName - Set the application name for the log messages
//
//	(c) 2023 Sam Caldwell.  MIT License
func (log *Logger) SetAppName(appName string) *Logger {
	if ValidAppName(&appName) {
		log.appName = appName
	} else {
		panic(errors.InvalidApplicationName)
	}
	return log
}
