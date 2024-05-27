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
type Logger[TGT LogTarget.TargetSet] struct {
	target  TGT
	level   LogLevel.Value
	appName string
	MsgId   string
}

// Configure - Configure the logger
//
//	(c) 2023 Sam Caldwell.  MIT License
func (log *Logger[T]) Configure(cfg configuration.Map[string, string]) *Logger[T] {
	if cfg == nil {
		panic("config cannot be nil")
	}
	log.target.Configure(cfg)
	return log
}

func (log *Logger[T]) Critical(message LogEvent.MessageValue) *Logger[T] {
	if log.level.Evaluate(LogLevel.Critical) {
		_ = log.target.Write(LogLevel.Critical, &message)
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
		_ = log.target.Write(&payload)
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
		_ = log.target.Write(&payload)
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
		_ = log.target.Write(&payload)
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
		_ = log.target.Write(&payload)
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
		_ = log.target.Write(&payload)
	}
	return log
}
