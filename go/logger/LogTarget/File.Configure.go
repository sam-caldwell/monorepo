package LogTarget

import (
	"github.com/sam-caldwell/monorepo/go/convert"
	"github.com/sam-caldwell/monorepo/go/misc/words"
	"os"
)

func (out File) Configure(cfg ConfigureTarget) {
	filename := cfg.ExpectOrPanic(words.FileName)
	permission := convert.StringToFileModeOrPanic(cfg.ExpectOrPanic(words.Permission))
	if err := out.file.Open(filename, os.O_CREATE|os.O_WRONLY|os.O_APPEND, permission); err != nil {
		panic(err)
	}
	cfg.rateLimit = convert.StringToUint(cfg.ExpectOrPanic(words.RateLimit))
}
