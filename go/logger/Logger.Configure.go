package logger

import "github.com/sam-caldwell/monorepo/go/logger/LogTarget"

// Configure - Configure the logger
func (log *Logger[T]) Configure(cfg LogTarget.ConfigureTarget) *Logger[T] {
	if cfg == nil {
		panic("config cannot be nil")
	}
	log.target.Configure(cfg)
	return log
}
