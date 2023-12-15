package psqlTrackerDb

import "github.com/google/uuid"

// TrackerIcon - This is a struct representing the Tracker DB Icon
type TrackerIcon struct {
	Id       uuid.UUID `json:"id"`
	Hash     string    `json:"hash"`
	MimeType string    `json:"mimeType"`
}
