package psqlTrackerDb

import "github.com/google/uuid"

// TrackerAvatar - This is a struct representing the Tracker DB Avatar
type TrackerAvatar struct {
	Id       uuid.UUID `json:"id"`
	Hash     string    `json:"hash"`
	MimeType string    `json:"mimeType"`
}
