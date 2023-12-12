package psqlTrackerDb

import "github.com/google/uuid"

// TrackerWorkflow - This is a struct representing the Tracker DB Workflow
type TrackerWorkflow struct {
	Id          uuid.UUID `json:"id"`
	Name        string    `json:"name"`
	IconId      uuid.UUID `json:"iconId"`
	OwnerId     uuid.UUID `json:"ownerId"`
	TeamId      uuid.UUID `json:"teamId"`
	Owner       string    `json:"owner"`
	Team        string    `json:"team"`
	Everyone    string    `json:"everyone"`
	Description string    `json:"description"`
}
