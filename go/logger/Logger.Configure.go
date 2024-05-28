package logger

import configuration "github.com/sam-caldwell/monorepo/go/configuration/Map"

// Configure - Configure the logger
//
//	(c) 2023 Sam Caldwell.  MIT License
func (log *Logger[T]) Configure(cfg configuration.Map[string, string]) *Logger[T] {
	if cfg == nil {
		panic("config cannot be nil")
	}
	log.target.Configure(cfg)
	return log
}
