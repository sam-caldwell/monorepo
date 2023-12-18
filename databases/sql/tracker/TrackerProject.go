package psqlTrackerDb

import (
	"github.com/google/uuid"
	"time"
)

type TrackerProject struct {
	Id          uuid.UUID `yaml:"id"`
	IconId      uuid.UUID `yaml:"iconId"`
	OwnerId     uuid.UUID `yaml:"ownerId"`
	TeamId      uuid.UUID `yaml:"teamId"`
	Created     time.Time `yaml:"created"`
	Name        string    `yaml:"name"`
	Owner       string    `yaml:"owner"`
	Team        string    `yaml:"team"`
	Everyone    string    `yaml:"everyone"`
	Description string    `yaml:"description"`
}
