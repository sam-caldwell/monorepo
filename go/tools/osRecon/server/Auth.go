package server

import (
	"net/http"
	"strings"
)

type HttpHandlerFunc func(http.ResponseWriter, *http.Request)

// Auth - Perform API Key authentication for the server listener
func (svr *Server) Auth(apiFunc HttpHandlerFunc) HttpHandlerFunc {

	const httpApiKeyHeaderName = "APIKEY"

	// Return the API handler
	return func(w http.ResponseWriter, r *http.Request) {

		if key := strings.TrimSpace(r.Header.Get(httpApiKeyHeaderName)); key == string(svr.apiKey) {

			// Invoke the payload function
			apiFunc(w, r)

		} else {

			// If APIKey is not correct, then reject as http/401 (unauthorized)
			w.WriteHeader(http.StatusUnauthorized)
			_, _ = w.Write([]byte("not authorized"))

		}
	}
}
