package server

import (
	"encoding/json"
	"github.com/google/uuid"
	"io"
	"net/http"
)

// getClientId - Parse the HTTP request for the clientId
func getClientId(r *http.Request, clientId *uuid.UUID) (err error) {
	var body []byte
	var request ApiCheckInRequest
	defer func() { _ = r.Body.Close() }()
	if body, err = io.ReadAll(r.Body); err != nil {
		return err
	}
	if err := json.Unmarshal(body, &request); err != nil {
		return err
	}
	*clientId = request.ClientId
	return err
}
