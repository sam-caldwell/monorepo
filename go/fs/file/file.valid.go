package file

import (
	"fmt"
	"regexp"
)

const (
	errInvalidFilePath = "invalid file/path"
)

// valid - validate the filepath input.
func (fp *File) valid(filePath *string) error {
	// Use the regular expression for file paths with file servers and traversal attempts
	const pattern = `^(?:(?:\.{0,1}\/(?:\w+\/)*|\w+:\/{2}(?:\w+:\w+@)?\w+(?:\.\w+)*(?::\d+)?\/(?:\w+\/)*)|(?:\w+:\/\/(?:\w+:\w+@)?[^\/]+\/(?:\w+\/)*))?([\w.-]+)$`

	validPathRegex := regexp.MustCompile(pattern)

	if validPathRegex.MatchString(*filePath) {
		return nil
	}
	return fmt.Errorf(errInvalidFilePath)
}
