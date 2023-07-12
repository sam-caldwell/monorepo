package projectmanifest

import (
	"fmt"
	"github.com/sam-caldwell/go/v2/projects/exit/errors"
	"strings"
)

func (manifest *Manifest) SetLanguage(language string) *Manifest {
	if manifest.err != nil {
		return manifest
	}
	switch thisLanguage := strings.TrimSpace(strings.ToLower(language)); thisLanguage {
	case "amd64Asm", "arm64Asm", "c", "cpp", "go", "node", "python", "rust", "typescript":
		manifest.Options.Language = language
	default:
		manifest.err = fmt.Errorf(errors.UnsupportedLanguage)
	}
	return manifest
}
