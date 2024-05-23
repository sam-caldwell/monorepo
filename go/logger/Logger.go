package logger

import (
	"github.com/sam-caldwell/monorepo/go/logger/LogEvent"
	"github.com/sam-caldwell/monorepo/go/logger/LogTarget"
)

// Logger - Top-level logging object
type Logger[TGT LogTarget.LogTarget, FMT LogEvent.LogFormat] struct {
	target TGT[FMT]
}
