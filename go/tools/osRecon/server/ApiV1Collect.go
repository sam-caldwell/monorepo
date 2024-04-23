package server

import (
	"github.com/sam-caldwell/monorepo/go/types"
	"net/http"
)

// ApiV1Collect - Provide payload for the Server Event Collection API endpoint (ApiV1Collect).
// Given a directory where this endpoint can write events (svr.eventCollectionDirectory)
func (svr *Server) ApiV1Collect(w http.ResponseWriter, r *http.Request) {

	//ToDo: create API endpoint to handle event collection
	/*
	 * the server needs a directory where it can write events. Each event
	 * is written to a client-specific subdirectory in a FIFO basis using
	 * an incrementing serial number.
	 */

}

func (svr *Server) WriteEvent(event types.Event) error {

	return nil
}
