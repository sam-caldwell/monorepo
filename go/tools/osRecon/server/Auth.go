package server

import (
	"net/http"
)

type HttpHandlerFunc func(http.ResponseWriter, *http.Request)

// Auth - Perform API Key authentication for the server listener
func (svr *Server) Auth(apiFunc HttpHandlerFunc) HttpHandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var authenticated bool
		//ToDo: inspect the request (r) for apiKey
		if authenticated {
			apiFunc(w, r)
			return
		} else {
			w.WriteHeader(http.StatusUnauthorized)
			_, _ = w.Write([]byte("not authorized"))
			return
		}
	}
}
