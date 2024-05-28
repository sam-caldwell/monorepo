package logger

import (
	configuration "github.com/sam-caldwell/monorepo/go/configuration/Map"
	"github.com/sam-caldwell/monorepo/go/convert"
	"github.com/sam-caldwell/monorepo/go/exit/errors"
	"github.com/sam-caldwell/monorepo/go/fs/file"
	"github.com/sam-caldwell/monorepo/go/misc/words"
	"os"
	"strings"
)

// Configure - Configure the logger
//
//	(c) 2023 Sam Caldwell.  MIT License
func (log *Logger) Configure(cfg configuration.Map[string, string]) *Logger {
	if cfg == nil {
		panic("config cannot be nil")
	}
	switch target := cfg.ExpectOrIgnore(words.Target); strings.TrimSpace(strings.ToLower(target)) {
	case words.File:
		var logFile file.File
		fileName := cfg.ExpectOrPanic(words.FileName)
		permission := convert.StringToFileModeOrPanic(cfg.ExpectOrPanic(words.Permission))
		if err := logFile.Open(fileName, os.O_CREATE|os.O_WRONLY|os.O_APPEND, permission); err != nil {
			panic("failed to open log file")
		}
		log.target = logFile.GetWriter()
	case words.Syslog:
		//ToDo: implement syslog network io.writer
		panic(errors.NotImplemented)
	case words.Http:
		//ToDo: implement Http network io.writer
		panic(errors.NotImplemented)
	case words.Stdout:
	default:
		log.target = os.Stdout
	}
	log.rateLimit = convert.StringToUint(cfg.ExpectOrIgnore(words.RateLimit))
	return log
}
