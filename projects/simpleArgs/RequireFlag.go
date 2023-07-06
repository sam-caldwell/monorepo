package simpleArgs

import "strings"

// RequireFlag - Return a boolean true or throw an error
func RequireFlag(name string) (bool, error) {
	if HasFlag(name) {
		return true, nil
	}
	return

	for _, option := range args.Options {

		if option == strings.TrimSpace(strings.ToLower(name)) {
			return true
		}
	}
	return false
}
