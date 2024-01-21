package file

import "fmt"

const (
	errMissingExtension = "missing extension"
)

// SetIfHasExtension - Set the given file/path name if the given name has one of the allowed extensions.
// Extensions must start with periods.
func (fp *File) SetIfHasExtension(name string, allowedExtensions []string) error {
	if err := fp.valid(&name); err != nil {
		*fp = ""
		return err
	}
	if !HasExtensions(name, allowedExtensions) {
		*fp = ""
		return fmt.Errorf(errMissingExtension)
	}
	*fp = File(name)
	return nil
}
