package server

import (
	"github.com/sam-caldwell/monorepo/go/fs/directory"
	"github.com/sam-caldwell/monorepo/go/tools/osRecon/cli"
)

// Configure - Configure the server from command-line
func (svr *Server) Configure() (err error) {

	var p *directory.Path

	svr.host = cli.GetHost()
	svr.port = cli.GetPort()
	svr.apiKey = cli.GetApiKey()

	p, err = cli.GetEventCollectorDir()
	if err != nil {
		return err
	}

	if err = svr.eventCollector.Init(p); err != nil {
		return err
	}

	if p, err = cli.GetQueryCollectorDir(); err != nil {
		return err
	}

	return svr.queryQueue.Init(p)

}
