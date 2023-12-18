package psqlTrackerDb

import "github.com/google/uuid"

// TrackerTicketType - This is a ticket type record
type TrackerTicketType struct {
	Id          uuid.UUID `json:"id"`
	Name        string    `json:"name"`
	IconId      uuid.UUID `json:"iconId"`
	WorkflowId  uuid.UUID `json:"workflowId"`
	Created     string    `json:"created"`
	Description string    `json:"description"`
}
