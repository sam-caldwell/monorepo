package projectmanifest

/*
 * projects/repotool/manifest/SetLanguage.go
 * (c) 2023 Sam Caldwell.  See LICENSE.txt
 *
 * This file defines SetLanguage() will set the project
 * language in the internal state
 */

import (
	"fmt"
	monorepo "github.com/sam-caldwell/go/v2/projects/__system__"
	"github.com/sam-caldwell/go/v2/projects/exit/errors"
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
