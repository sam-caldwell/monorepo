package server

import (
	"github.com/sam-caldwell/monorepo/go/ansi"
	"github.com/sam-caldwell/monorepo/go/exit"
	"github.com/sam-caldwell/monorepo/go/fs/directory"
	"github.com/sam-caldwell/monorepo/go/simpleArgs"
)

func (svr *Server) Configure() error {
	var p *directory.Path
	argEventCollectorDir, err := simpleArgs.GetOptionValue("--eventCollectorDir")
	if err != nil {
		ansi.Red().
			Printf("Parsing error (%s): %v", "--eventCollectorDir", err.Error()).
			LF().
			Fatal(exit.GeneralError).
			Reset()
	}
	p, svr.err = directory.NewPath(argEventCollectorDir, true)
	svr.err = svr.eventCollector.Init(p)
	return svr.err
}
