package projectFilter

import (
	"fmt"
	keyvalue "github.com/sam-caldwell/go/v2/projects/KeyValue"
	"github.com/sam-caldwell/go/v2/projects/exit/errors"
	"github.com/sam-caldwell/go/v2/projects/fs/file"
	"path/filepath"
)

func FilterThisPath(filter Filter, baseDirectory string) bool {
	manifest, err := LoadProjectManifest(baseDirectory)
	if err != nil {
		return true //We filter any project that lacks a (properly formatted) MANIFEST.txt file.
	}
	hasFlag := func(name string) bool {
		_, found := manifest.FindKey("project.disabled")
		return found
	}

	return (filter.HasFilter(HideDisabled) && hasFlag("project.disabled")) ||
		(filter.HasFilter(HideEnabled) && !hasFlag("project.disabled")) ||
		(filter.HasFilter(Windows) && hasFlag("windows.disabled")) ||
		(filter.HasFilter(Linux) && hasFlag("linux.disabled")) ||
		(filter.HasFilter(Darwin) && hasFlag("darwin.disabled")) ||
		(filter.HasFilter(LanguageAsmAmd64) && FilterOnLanguage(manifest, LanguageAsmAmd64)) ||
		(filter.HasFilter(LanguageAsmArm64) && FilterOnLanguage(manifest, LanguageAsmArm64)) ||
		(filter.HasFilter(LanguageC) && FilterOnLanguage(manifest, LanguageC)) ||
		(filter.HasFilter(LanguageCpp) && FilterOnLanguage(manifest, LanguageCpp)) ||
		(filter.HasFilter(LanguageGo) && FilterOnLanguage(manifest, LanguageGo)) ||
		(filter.HasFilter(LanguageNodeJs) && FilterOnLanguage(manifest, LanguageNodeJs)) ||
		(filter.HasFilter(LanguagePython) && FilterOnLanguage(manifest, LanguagePython)) ||
		(filter.HasFilter(LanguageRust) && FilterOnLanguage(manifest, LanguageRust))
}

func LoadProjectManifest(projectPath string) (kv keyvalue.KeyValue, err error) {
	manifest := filepath.Join(projectPath, "MANIFEST.txt")
	if !file.Exists(manifest) {
		return kv, fmt.Errorf(errors.NotFound)
	}
	if err := kv.FromFile(manifest, ":", "\n"); err != nil {
		return kv, err
	}
	return kv, err
}
