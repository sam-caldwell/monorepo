package systemrecon

import "os/user"

// GetCurrentUserName - Return current username
func GetCurrentUserName() (string, error) {
	currentUser, err := user.Current()
	return currentUser.Name, err
}
