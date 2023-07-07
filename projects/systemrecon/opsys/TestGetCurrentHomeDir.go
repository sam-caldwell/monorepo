package systemrecon

import (
	"os/user"
)

// GetCurrentHomeDir - Return the current user's home directory
func GetCurrentHomeDir() (string, error) {
	currentUser, err := user.Current()
	return currentUser.HomeDir, err
}
