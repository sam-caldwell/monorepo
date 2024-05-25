package LogTarget

import (
	"github.com/sam-caldwell/monorepo/go/ansi"
	"github.com/sam-caldwell/monorepo/go/fs/file"
)

// Write - Write a formatted log string to stdout.
//
//	(c) 2023 Sam Caldwell.  MIT License
func (out File) Write(p *[]byte) {
	if out.file == nil {
		out.file = file.Open(out.config.fileName)
		panic("file logger not configured")
	}
	ansi.Println(string(*p))
}
