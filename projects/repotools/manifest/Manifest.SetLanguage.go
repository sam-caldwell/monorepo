package projectmanifest

import (
	"fmt"
	monorepo "github.com/sam-caldwell/go/v2/projects/__system__"
	"github.com/sam-caldwell/go/v2/projects/exit/errors"
)

func (manifest *Manifest) SetLanguage(language string) *Manifest {
	if manifest.err != nil {
		return manifest
	}

	for _, supported := range monorepo.GetSupportedLanguages() {
		if language == supported {
			manifest.Options.Language = language
			return manifest
		}
	}
	manifest.err = fmt.Errorf(errors.UnsupportedLanguage)
	return manifest
}
