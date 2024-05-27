package LogTarget

import (
	"github.com/sam-caldwell/monorepo/go/ansi"
	"github.com/sam-caldwell/monorepo/go/logger/LogEvent"
	"github.com/sam-caldwell/monorepo/go/logger/LogLevel"
)

// Write - Write a formatted log string to stdout.
func (out *StdoutTarget) Write(messagePriority LogLevel.Value, message *LogEvent.MessageValue) error {
	var event LogEvent.RFC5424Message
	event.Create(messagePriority, &out.appName, &out.msgId, message)
	if payload, err := event.ToJsonString(); err != nil {
		return err
	} else {
		ansi.Println(payload)
	}
	return nil
}
