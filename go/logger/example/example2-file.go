package main

import (
	"github.com/sam-caldwell/monorepo/go/ansi"
	"github.com/sam-caldwell/monorepo/go/logger"
	"github.com/sam-caldwell/monorepo/go/logger/LogLevel"
	"github.com/sam-caldwell/monorepo/go/logger/LogTarget"
)

func main() {
	ansi.White().Println("Test starting...")
	// Declare the logger and specify the output target (e.g. stdout, file, http,...)
	var log logger.Logger[LogTarget.File]
	// Configure the log level and other parameters
	log.Configure(LogTarget.ConfigureTarget{
		"filename":   "/tmp/test.log",
		"permission": "0600",
	}).
		SetLevel(LogLevel.Debug).
		Debug("debug message (expect blue text)").
		Info("info message (expect blue text)").
		Warning("warning message (expect yellow text)").
		Error("error message (expect red text)").
		Critical("critical message (expect red text)")
	log.SetLevel(LogLevel.Info).
		Debug("debug message (expect blue text)").
		Info("info message (expect blue text)").
		Warning("warning message (expect yellow text)").
		Error("error message (expect red text)").
		Critical("critical message (expect red text)")
	log.SetLevel(LogLevel.Warning).
		Debug("debug message (expect blue text)").
		Info("info message (expect blue text)").
		Warning("warning message (expect yellow text)").
		Error("error message (expect red text)").
		Critical("critical message (expect red text)")
	log.SetLevel(LogLevel.Error).
		Debug("debug message (expect blue text)").
		Info("info message (expect blue text)").
		Warning("warning message (expect yellow text)").
		Error("error message (expect red text)").
		Critical("critical message (expect red text)")
	log.SetLevel(LogLevel.Critical).
		Debug("debug message (expect blue text)").
		Info("info message (expect blue text)").
		Warning("warning message (expect yellow text)").
		Error("error message (expect red text)").
		Critical("critical message (expect red text)")
	ansi.White().Println("test finishing").Reset()
}
