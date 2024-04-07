package server

import (
	"fmt"
	"net/http"
	"time"
)

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
