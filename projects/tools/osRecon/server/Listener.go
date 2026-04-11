package server

import (
	"fmt"
	"net/http"
)

// Listener - Listen on either http/1.1 or http/2 for API requests.
func (svr *Server) Listener() (err error) {

	go func() {
		//
		// setup our API endpoints
		//
		http.HandleFunc(ApiV1CheckIn, svr.Auth(svr.ApiV1CheckIn))
		http.HandleFunc(ApiV1Collect, svr.Auth(svr.ApiV1Collect))
		//
		// ToDo: Implement TLS encryption
		//
		err = http.ListenAndServe(fmt.Sprintf("%s:%d", svr.host, svr.port), nil)
	}()

	return err

}
