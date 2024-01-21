package file

import (
	"fmt"
	"github.com/sam-caldwell/monorepo/go/misc/words"
	"strings"
)

const (
	errMissingExtension = "missing extension"
)

// SetWithExtension - Set the given file/path name to the internal state (no validation), verifying the .iso extension
func (fp *File) SetWithExtension(name string, allowedExtensions []string) error {
	if err := fp.valid(&name); err != nil {
		for _, ext := range allowedExtensions {
			found := false
			if !strings.HasPrefix(ext, words.Period) {
				ext = fmt.Sprintf("%s%s", words.Period, ext)
			}
			if ok := strings.HasSuffix(string(*fp), ext); !ok {
				found = found || ok
			}
			if !found {
				return fmt.Errorf(errMissingExtension)
			}
		}
		*fp = ""
		return err
	}
	*fp = File(name)
	return nil
}
