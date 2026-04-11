package logger

import (
	configuration "github.com/sam-caldwell/monorepo/projects/configuration/Map"
	"github.com/sam-caldwell/monorepo/projects/convert"
	"github.com/sam-caldwell/monorepo/projects/exit/errors"
	"github.com/sam-caldwell/monorepo/projects/fs/file"
	"github.com/sam-caldwell/monorepo/projects/misc/words"
	"os"
)

// ConfigureFile - Configure a log file
//
//	(c) 2023 Sam Caldwell.  MIT License
func (log *Logger) ConfigureFile(cfg *configuration.Map[string, string]) {
	var logFile file.File
	fileName := cfg.ExpectOrPanic(words.FileName)
	permission := convert.StringToFileModeOrPanic(cfg.ExpectOrPanic(words.Permission))
	if err := logFile.Open(fileName, os.O_CREATE|os.O_WRONLY|os.O_APPEND, permission); err != nil {
		panic(errors.FailedToOpenLogFile)
	}
	log.useColor = false
	log.target = logFile.GetWriter()
}
