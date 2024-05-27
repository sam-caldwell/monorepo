package file

import (
	"fmt"
	"regexp"
)

const (
	errInvalidFilePath = "invalid file/path"
)

// valid - validate the filepath input.
//
//	     Note: any error returned indicates a validation issue.  Nil returns indicate valid input.
//
//		    Directory traversal should not be allowed.
//
//			(c) 2023 Sam Caldwell.  MIT License
func (fp *File) valid(filePath *string) error {
	// Use the regular expression for file paths with file servers and traversal attempts
	const pattern = `` +
		`^(?:(?:\.{0,1}\/(?:\w+\/)*|\w+:\/{2}(?:\w+:\w+@)?\w+(?:\.\w+)*(?::\d+)?\/(?:\w+\/)*)` +
		`|(?:\w+:\/\/(?:\w+:\w+@)?[^\/]+\/(?:\w+\/)*))?([\w.-]+)$`

	validPathRegex := regexp.MustCompile(pattern)

	if validPathRegex.MatchString(*filePath) {
		return nil
	}
	return fmt.Errorf(errInvalidFilePath)
}
