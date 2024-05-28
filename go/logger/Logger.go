package logger

import (
	"github.com/sam-caldwell/monorepo/go/logger/LogTarget"
)

// Logger - Top-level logging object
//
//	(c) 2023 Sam Caldwell.  MIT License
type Logger[TGT LogTarget.TargetSet] struct {
	target TGT
}
