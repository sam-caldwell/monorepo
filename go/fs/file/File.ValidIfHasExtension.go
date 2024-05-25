package file

import "fmt"

const (
	errMissingExtension = "missing extension"
)

// ValidIfHasExtension - Return nil if the file name is valid and has an allowed extension.
//
//	     Note: any error returned indicates a validation issue.  Nil returns indicate valid input
//
//			(c) 2023 Sam Caldwell.  MIT License
func (fp *File) ValidIfHasExtension(name string, allowedExtensions []string) error {
	if err := fp.valid(&name); err != nil {
		return err
	}
	if !HasExtensions(name, allowedExtensions) {
		return fmt.Errorf(errMissingExtension)
	}
	return nil
}
