package server

import "net/http"

func (svr *Server) ApiV1Collect(w http.ResponseWriter, r *http.Request) {

	//ToDo: create API endpoint to handle event collection
	/*
	 * the server needs a directory where it can write events. Each event
	 * is written to a client-specific subdirectory in a FIFO basis using
	 * an incrementing serial number.
	 */

}
