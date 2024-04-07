package server

import "net/http"

func (svr *Server) ApiV1CheckIn(w http.ResponseWriter, r *http.Request) {

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
