package connector

import (
	"github.com/sam-caldwell/monorepo/go/ansi"
	"github.com/sam-caldwell/monorepo/go/types"
)

/*
 * ToDo: we need to wrap the events with meta data about the machine.
 *       we might do this with an event.Wrap() method that will return
 *       a wrapper struct containing metadata and the original event.
 */

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
