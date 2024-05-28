package LogTarget

import (
	"github.com/sam-caldwell/monorepo/go/logger/LogEvent"
	logLevel "github.com/sam-caldwell/monorepo/go/logger/LogLevel"
)

// TargetSet - A type of all LogTargetTypes
//
//	(c) 2023 Sam Caldwell.  MIT License
type TargetSet interface {
	StdoutTarget | FileTarget | HttpTarget | SyslogTarget
	SetAppName(*string)
	Write(value logLevel.Value, messageValue LogEvent.MessageValue) error
}
