package types

import "github.com/google/uuid"

// ClientId - a uuid-based client identifier
type ClientId uuid.UUID

// String - Return the string form of the ClientId
func (c *ClientId) String() string {
	return uuid.UUID(*c).String()
}

// FromString - Given a string, parse it as a UUID and store as ClientId
func (c *ClientId) FromString(s string) (err error) {
	var u uuid.UUID
	if u, err = uuid.Parse(s); err == nil {
		*c = ClientId(u)
	}
	return err
}
