package psqlTrackerDb

import "github.com/google/uuid"

// TrackerEntity - This is a struct representing the Tracker DB Entity
type TrackerEntity struct {
	Id      uuid.UUID `json:"id"`
	Type    string    `json:"type"`
	Created string    `json:"created"`
	Context string    `json:"context"`
}
