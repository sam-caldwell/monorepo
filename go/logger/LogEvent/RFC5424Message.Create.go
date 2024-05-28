package LogEvent

import (
	"github.com/sam-caldwell/monorepo/go/logger/LogLevel"
	"github.com/sam-caldwell/monorepo/go/version"
	"os"
	"time"
)

// Create - Create and configure an RFC5424 message object
//
//	(c) 2023 Sam Caldwell.  MIT License
func (e *RFC5424Message) Create(messagePriority LogLevel.Value, appName, msgId *string, message *MessageValue) *RFC5424Message {
	var err error
	e.Priority = uint(messagePriority)
	e.Version = version.Version
	e.Timestamp = time.Now()
	if e.Hostname, err = os.Hostname(); err != nil {
		e.Hostname = "not_available"
	}
	e.AppName = *appName
	e.MsgID = *msgId
	e.Message = *message
	return e
}
