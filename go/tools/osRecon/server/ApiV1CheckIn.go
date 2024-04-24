package server

import (
    "github.com/google/uuid"
    "log"
    "net/http"
)

// ApiV1CheckIn - Provide payload for the Server Check-in API endpoint (ApiV1CheckIn).
func (svr *Server) ApiV1CheckIn(w http.ResponseWriter, r *http.Request) {
	const (
		queryFetchCount uint16 = 5
	)
	if r.Method == http.MethodGet {
		var clientId uuid.UUID

		if err := getClientId(r, &clientId); err != nil {
			w.WriteHeader(http.StatusBadRequest)
			_, _ = w.Write([]byte("Bad Request"))
			return
		}

		if err := svr.queryQueue.FetchQueryList(clientId, queryFetchCount); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			_, _ = w.Write([]byte("QueryFetchError()"))
			log.Printf("QueryFetchError: %v", err)
		}
		//ToDo: log the heartbeat with the clientId
	}
	//ToDo: create API endpoint to handle healthcheck / query queue polling
}
