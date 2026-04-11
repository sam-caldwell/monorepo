package types

import (
	"fmt"
	"regexp"
)

// ApiKey - Api Key for authenticating with server
type ApiKey []byte

func (o *ApiKey) FromString(v *string) error {
	pattern := regexp.MustCompile(`^[A-Za-z0-9+/]{64}$`)
	if pattern.MatchString(*v) {
		return nil
	}
	return fmt.Errorf("invalid apiKey")
}
