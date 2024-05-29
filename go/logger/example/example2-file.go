package main

import (
	"github.com/sam-caldwell/monorepo/go/ansi"
	configuration "github.com/sam-caldwell/monorepo/go/configuration/Map"
	"github.com/sam-caldwell/monorepo/go/logger"
	"github.com/sam-caldwell/monorepo/go/logger/LogLevel"
)

// main - File-based logging solution
//
//	This log implementation writes logs to the file.
//
//	(c) 2023 Sam Caldwell.  MIT License
func main() {
	LoggerExampleFile()
}

func LoggerExampleFile() {
	ansi.White().Println("Test starting...")

	var log logger.Logger
	// Configure the log level and other parameters
	log.Configure(&(configuration.Map[string, string]{
		"appName":    "example2",
		"msgId":      "99",
		"target":     "file",
		"filename":   "/tmp/test.log",
		"permission": "0600", //file permission
		"rateLimit":  "100",
	})).
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
