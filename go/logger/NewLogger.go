package logger

import (
	"github.com/sam-caldwell/monorepo/go/logger/LogEvent"
	"github.com/sam-caldwell/monorepo/go/logger/LogTarget"
)

// NewLogger - Create and return a new logger
func NewLogger[TGT LogTarget.LogTarget, FMT LogEvent.LogFormat]() *Logger[TGT, FMT] {
	var log Logger[TGT, FMT]
	return &log
}
