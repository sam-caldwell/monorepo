package types

import "github.com/google/uuid"

type ClientId uuid.UUID

func (c *ClientId) String() string {
	return uuid.UUID(*c).String()
}
