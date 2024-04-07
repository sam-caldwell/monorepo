package server

import (
	"fmt"
	"net/http"
	"time"
)

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

//ToDo: create API endpoint to handle event collection
/*
 * the server needs a directory where it can write events. Each event
 * is written to a client-specific subdirectory in a FIFO basis using
 * an incrementing serial number.
 */

// Listener - Listen on either http/1.1 or http/2 for API requests.
func (svr *Server) Listener() *Server {
	go func() {
		http.HandleFunc("/api/v1/checkIn", svr.Auth(svr.ApiV1CheckIn))
		http.HandleFunc("/api/v1/collect", svr.Auth(svr.ApiV1Collect))
		//ToDo: implement TLS encryption
		svr.err = http.ListenAndServe(fmt.Sprintf("%s:%d", svr.host, svr.port), nil)
	}()
	time.Sleep(1 * time.Second)
	return svr
}
