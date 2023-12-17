package psqlTrackerDb

import "github.com/google/uuid"

// TrackerTicketType - This is a ticket type record
type TrackerTicketType struct {
	Id          uuid.UUID `yaml:"id"`
	Name        string    `yaml:"name"`
	IconId      uuid.UUID `yaml:"iconId"`
	WorkflowId  uuid.UUID `yaml:"workflowId"`
	Created     string    `yaml:"created"`
	Description string    `yaml:"description"`
}
