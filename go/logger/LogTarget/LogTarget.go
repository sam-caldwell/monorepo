package LogTarget

import (
	configuration "github.com/sam-caldwell/monorepo/go/configuration/Map"
)

// LogTarget - The target to which logs will be written
//
//	(c) 2023 Sam Caldwell.  MIT License
type LogTarget interface {
	Configure(cfg configuration.Map[string, string])
}
