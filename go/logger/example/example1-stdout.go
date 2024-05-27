package main

import (
	"github.com/sam-caldwell/monorepo/go/ansi"
	"github.com/sam-caldwell/monorepo/go/logger"
	"github.com/sam-caldwell/monorepo/go/logger/LogLevel"
	"github.com/sam-caldwell/monorepo/go/logger/LogTarget"
)

// main - Simple stdout logging implementation
//
//	 This log implementation is the minimum configuration possible
//	 and uses stdout as the log target.
//
//		(c) 2023 Sam Caldwell.  MIT License
func main() {
	ansi.White().Println("Test starting...")
	// Declare the logger and specify the output target (e.g. stdout, file, http,...)
	var log logger.Logger[LogTarget.StdoutTarget]
	// Configure the log level and other parameters
	log.SetLevel(LogLevel.Debug).
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
