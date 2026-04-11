package Atlassian

import (
	"fmt"
)

// SetUsername - Set the username
func (client *Client) SetUsername(name *string) error {

	if name == nil || *name == "" {

		return fmt.Errorf("missing username (use --username)")

	}

	client.username = *name

	return nil

}
