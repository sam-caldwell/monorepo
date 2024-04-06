package server

import (
	"github.com/sam-caldwell/monorepo/go/ansi"
	"github.com/sam-caldwell/monorepo/go/types"
)

// SendEvent - Send an event object to the server
func (svr *Server) SendEvent(event types.Event) *Server {
	//Only attempt to send if error state is clear
	if svr.err == nil {
		//ToDo: Send the event to the server
		ansi.Cyan().
			Printf("Debug (SendEvent):%v", event.ToString()).
			LF().
			Reset()
		//ToDo: set err state if error occurred in send.
		svr.err = nil
	}

	return svr
}
