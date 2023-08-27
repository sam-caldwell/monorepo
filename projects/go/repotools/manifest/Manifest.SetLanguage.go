package projectmanifest

/*
 * projects/repotool/manifest/Manifest.SetLanguage.go
 * (c) 2023 Sam Caldwell.  See LICENSE.txt
 *
 * This file defines SetLanguage() will set the project
 * language in the internal state
 */

import (
	"fmt"
	"github.com/sam-caldwell/go/v2/projects/go/__system__"
	"github.com/sam-caldwell/go/v2/projects/go/exit/errors"
)

// SetLanguage - set project programming language
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
