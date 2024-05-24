package logger

import (
	"github.com/sam-caldwell/monorepo/go/logger/LogEvent"
	"github.com/sam-caldwell/monorepo/go/logger/LogLevel"
	"github.com/sam-caldwell/monorepo/go/version"
	"os"
	"time"
)

func (log *Logger[T, F]) Fatal(message LogEvent.MessageValue) *Logger[T, F] {
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
			Priority:  uint(log.level),
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
		log.target.Write(&payload)
		log.target.Flush()
	}
	return log
}
