package LogTarget

import (
	configuration "github.com/sam-caldwell/monorepo/go/configuration/Map"
	"github.com/sam-caldwell/monorepo/go/convert"
	"github.com/sam-caldwell/monorepo/go/misc/words"
	"os"
)

// Configure - Configure LogTarget file
//
//	(c) 2023 Sam Caldwell.  MIT License
func (out FileTarget) Configure(cfg configuration.Map[string, string]) {
	filename := cfg.ExpectOrPanic(words.FileName)
	permission := convert.StringToFileModeOrPanic(cfg.ExpectOrPanic(words.Permission))
	if err := out.file.Open(filename, os.O_CREATE|os.O_WRONLY|os.O_APPEND, permission); err != nil {
		panic(err)
	}
	out.rateLimit = convert.StringToUint(cfg.ExpectOrPanic(words.RateLimit))
}
