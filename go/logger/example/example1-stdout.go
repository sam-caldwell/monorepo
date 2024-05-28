package main

import (
	"github.com/sam-caldwell/monorepo/go/ansi"
	"github.com/sam-caldwell/monorepo/go/logger"
	"github.com/sam-caldwell/monorepo/go/logger/LogLevel"
)

// main - Simple stdout logging implementation
//
//	 This log implementation is the minimum configuration possible
//	 and uses stdout as the log target.
//
//		(c) 2023 Sam Caldwell.  MIT License
func main() {
	LoggerExampleStdout()
}

func LoggerExampleStdout() {
	ansi.White().Println("Test starting...").LF()

	var log logger.Logger
	log.DefaultConfiguration().
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
	ansi.White().LF().Println("test finishing").Reset()
}
