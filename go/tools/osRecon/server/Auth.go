package server

import (
	"net/http"
)

type HttpHandlerFunc func(http.ResponseWriter, *http.Request)

// Auth - Perform API Key authentication for the server listener
func (svr *Server) Auth(apiFunc HttpHandlerFunc) HttpHandlerFunc {
	//ToDo: inspect the API Key and request metadata
	var authenticated bool

	if authenticated {
		return apiFunc
	} else {
		return func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusUnauthorized)
			_, svr.err = w.Write([]byte("not authorized"))
		}
	}
}
