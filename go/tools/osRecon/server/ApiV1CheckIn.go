package server

import (
	"net/http"
)

// ApiV1CheckIn - Provide payload for the Server Check-in API endpoint (ApiV1CheckIn).
func (svr *Server) ApiV1CheckIn(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		//ToDo: log the checkin from the client as its heartbeat.

	}
	//ToDo: create API endpoint to handle healthcheck / query queue polling
	/*
	 * The server needs a directory where queries can be added (one per file)
	 * and a client can request the queries pending for it...
	 *
	 * <queue_directory>/<clientUUID>/<queue_record_filename>.tql
	 *
	 * queue_record_filenames should be numeric to allow them to be
	 * executed in order.
	 */
}
