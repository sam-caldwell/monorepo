package logger

import (
	"github.com/sam-caldwell/monorepo/go/logger/LogEvent"
	logLevel "github.com/sam-caldwell/monorepo/go/logger/LogLevel"
	"github.com/sam-caldwell/monorepo/go/logger/LogTarget"
)

// Logger - Top-level logging object
type Logger[TGT LogTarget.LogTarget, FMT LogEvent.LogFormat] struct {
	target TGT[FMT]
	level  logLevel.Value
}
