package main

import (
	"github.com/sam-caldwell/monorepo/go/logger"
	"github.com/sam-caldwell/monorepo/go/logger/LogEvent"
	"github.com/sam-caldwell/monorepo/go/logger/LogTarget"
	"github.com/sam-caldwell/monorepo/go/logger/logLevel"
)

func main() {
	var log logger.Logger[LogTarget.Stdout[LogEvent.RFC5424Message], LogEvent.RFC5424Message]
	log.SetLevel(logLevel.Debug)
}
