package logger

import (
	logLevel "github.com/sam-caldwell/monorepo/go/logger/LogLevel"
	"github.com/sam-caldwell/monorepo/go/logger/LogTarget"
)

// Logger - Top-level logging object
type Logger[TGT LogTarget.LogTarget] struct {
	target  TGT
	level   logLevel.Value
	appName string
	MsgId   string
}
