package psqlTrackerDb

import "github.com/google/uuid"

// TrackerTeam - This is a struct representing the Tracker DB Team
type TrackerTeam struct {
	Id          uuid.UUID `yaml:"id"`
	IconId      uuid.UUID `yaml:"iconId"`
	OwnerId     uuid.UUID `yaml:"ownerId"`
	Name        string    `yaml:"name"`
	Owner       string    `yaml:"owner"`
	Team        string    `yaml:"team"`
	Everyone    string    `yaml:"everyone"`
	Description string    `yaml:"description"`
}
