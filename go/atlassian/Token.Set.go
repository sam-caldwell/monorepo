package Atlassian

import (
	"fmt"
	"regexp"
)

// Set - set a token (after validating it)
func (o *Token) Set(v *string) (err error) {

	if (v == nil) || (*v == "") {
		return fmt.Errorf("empty or nil apiKey")
	}

	pattern := regexp.MustCompile(atlassianTokenRegex)

	if !pattern.MatchString(*v) {

		err = fmt.Errorf(errInvalidAtlassianToken)

	} else {

		*o = Token(*v)

	}

	return err

}
