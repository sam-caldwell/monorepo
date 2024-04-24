package server

import "github.com/google/uuid"

// ApiCheckInRequest - Representation of an API V1 Checkin request
type ApiCheckInRequest struct {
	ClientId uuid.UUID `json:"ClientId"`
}
