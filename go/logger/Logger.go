package logger

import (
    "github.com/sam-caldwell/monorepo/go/configuration/Map"
    "github.com/sam-caldwell/monorepo/go/logger/LogEvent"
	"github.com/sam-caldwell/monorepo/go/logger/LogLevel"
	"github.com/sam-caldwell/monorepo/go/logger/LogTarget"
	"github.com/sam-caldwell/monorepo/go/version"
	"os"
	"time"
)

// Logger - Top-level logging object
//
//	(c) 2023 Sam Caldwell.  MIT License
type Logger[TGT LogTarget.LogTarget] struct {
	target  TGT
	level   LogLevel.Value
	appName string
	MsgId   string
}

// Configure - Configure the logger
//
//	(c) 2023 Sam Caldwell.  MIT License
func (log *Logger[T]) Configure(cfg configuration.Map[string,string]) *Logger[T] {
	if cfg == nil {
		panic("config cannot be nil")
	}
	log.target.Configure(cfg)
	return log
}

// SetLevel - Define the log level
//
//	(c) 2023 Sam Caldwell.  MIT License
func (log *Logger[T]) SetLevel(level LogLevel.Value) *Logger[T] {
	log.level = level
	return log
}

// SetAppName - Set the application name for the log messages
//
//	(c) 2023 Sam Caldwell.  MIT License
func (log *Logger[T]) SetAppName(appName string) *Logger[T] {
	//ToDo: sanitize app name
	log.appName = appName
	return log
}

func (log *Logger[T]) Critical(message LogEvent.MessageValue) *Logger[T] {
	if log.level.Evaluate(LogLevel.Critical) {
		var (
			err      error
			hostname string
			payload  []byte
		)
		hostname, err = os.Hostname()
		if err != nil {
			hostname = "not_available"
		}
		payload, err = LogEvent.RFC5424Message{
			Priority:  uint(LogLevel.Critical),
			Version:   version.Version,
			Timestamp: time.Now(),
			Hostname:  hostname,
			AppName:   log.appName,
			ProcID:    os.Getpid(),
			MsgID:     log.MsgId,
			Message:   message,
		}.ToJson()
		if err != nil {
			panic("log message serialization error")
		}
		log.target.SetLevel(LogLevel.Critical)
		log.target.Write(&payload)
	}
	return log
}

func (log *Logger[T]) Error(message LogEvent.MessageValue) *Logger[T] {
	if log.level.Evaluate(LogLevel.Error) {
		var (
			err      error
			hostname string
			payload  []byte
		)
		hostname, err = os.Hostname()
		if err != nil {
			hostname = "not_available"
		}
		payload, err = LogEvent.RFC5424Message{
			Priority:  uint(LogLevel.Error),
			Version:   version.Version,
			Timestamp: time.Now(),
			Hostname:  hostname,
			AppName:   log.appName,
			ProcID:    os.Getpid(),
			MsgID:     log.MsgId,
			Message:   message,
		}.ToJson()
		if err != nil {
			panic("log message serialization error")
		}
		log.target.SetLevel(LogLevel.Error)
		log.target.Write(&payload)
	}
	return log
}

func (log *Logger[T]) Fatal(message LogEvent.MessageValue) *Logger[T] {
	if log.level.Evaluate(LogLevel.Fatal) {
		var (
			err      error
			hostname string
			payload  []byte
		)
		hostname, err = os.Hostname()
		if err != nil {
			hostname = "not_available"
		}
		payload, err = LogEvent.RFC5424Message{
			Priority:  uint(LogLevel.Fatal),
			Version:   version.Version,
			Timestamp: time.Now(),
			Hostname:  hostname,
			AppName:   log.appName,
			ProcID:    os.Getpid(),
			MsgID:     log.MsgId,
			Message:   message,
		}.ToJson()
		if err != nil {
			panic("log message serialization error")
		}
		log.target.SetLevel(LogLevel.Fatal)
		log.target.Write(&payload)
	}
	defer func() {
		os.Exit(1)
	}()
	return log
}

func (log *Logger[T]) Warning(message LogEvent.MessageValue) *Logger[T] {
	if log.level.Evaluate(LogLevel.Warning) {
		var (
			err      error
			hostname string
			payload  []byte
		)
		hostname, err = os.Hostname()
		if err != nil {
			hostname = "not_available"
		}
		payload, err = LogEvent.RFC5424Message{
			Priority:  uint(LogLevel.Warning),
			Version:   version.Version,
			Timestamp: time.Now(),
			Hostname:  hostname,
			AppName:   log.appName,
			ProcID:    os.Getpid(),
			MsgID:     log.MsgId,
			Message:   message,
		}.ToJson()
		if err != nil {
			panic("log message serialization error")
		}
		log.target.SetLevel(LogLevel.Warning)
		log.target.Write(&payload)
	}
	return log
}

func (log *Logger[T]) Info(message LogEvent.MessageValue) *Logger[T] {
	if log.level.Evaluate(LogLevel.Info) {
		var (
			err      error
			hostname string
			payload  []byte
		)
		hostname, err = os.Hostname()
		if err != nil {
			hostname = "not_available"
		}
		payload, err = LogEvent.RFC5424Message{
			Priority:  uint(LogLevel.Info),
			Version:   version.Version,
			Timestamp: time.Now(),
			Hostname:  hostname,
			AppName:   log.appName,
			ProcID:    os.Getpid(),
			MsgID:     log.MsgId,
			Message:   message,
		}.ToJson()
		if err != nil {
			panic("log message serialization error")
		}
		log.target.SetLevel(LogLevel.Info)
		log.target.Write(&payload)
	}
	return log
}

func (log *Logger[T]) Debug(message LogEvent.MessageValue) *Logger[T] {
	if log.level.Evaluate(LogLevel.Debug) {
		var (
			err      error
			hostname string
			payload  []byte
		)
		hostname, err = os.Hostname()
		if err != nil {
			hostname = "not_available"
		}
		payload, err = LogEvent.RFC5424Message{
			Priority:  uint(LogLevel.Debug),
			Version:   version.Version,
			Timestamp: time.Now(),
			Hostname:  hostname,
			AppName:   log.appName,
			ProcID:    os.Getpid(),
			MsgID:     log.MsgId,
			Message:   message,
		}.ToJson()
		if err != nil {
			panic("log message serialization error")
		}
		log.target.SetLevel(LogLevel.Debug)
		log.target.Write(&payload)
	}
	return log
}
