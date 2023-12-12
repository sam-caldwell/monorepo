package psqlTrackerDb

import "github.com/google/uuid"

// TrackerUser - This is a struct representing the Tracker DB User
type TrackerUser struct {
	Id          uuid.UUID `yaml:"id"`
	FirstName   string    `yaml:"firstName"`
	LastName    string    `yaml:"lastName"`
	AvatarId    uuid.UUID `yaml:"avatarId"`
	Email       string    `yaml:"email"`
	PhoneNumber string    `yaml:"phoneNumber"`
	Description string    `yaml:"description"`
}
