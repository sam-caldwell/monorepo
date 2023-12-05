package monorepo

import (
	"fmt"
	"github.com/sam-caldwell/monorepo/go/fs/directory"
	"github.com/sam-caldwell/monorepo/go/fs/file"
	"github.com/sam-caldwell/monorepo/go/misc/words"
	"path/filepath"
	"strings"
)

func GetProjectList(criteria *string) (list []string, err error) {
	rootDirectory := directory.GetCurrent()
	var searchSpaces []string
	if strings.TrimSpace(*criteria) == "" {
		return list, fmt.Errorf("invalid project specifier (provide comma-delimited list or 'all')")
	}
	if strings.TrimSpace(strings.ToLower(*criteria)) == "all" {
		searchSpaces = []string{
			"asm", "containers", "cpp", "databases/graph",
			"databases/sql", "databases/tsdb", "go", "js", "python"}
	} else {
		searchSpaces = strings.Split(*criteria, words.Comma)
	}
	for _, projectDirectory := range searchSpaces {
		if !directory.Exists(filepath.Join(rootDirectory, projectDirectory)) {
			err = fmt.Errorf("directory (%s) not found", projectDirectory)
			break
		}
		manifest := fmt.Sprintf("%s/Manifest.yaml", projectDirectory)
		if !file.Exists(manifest) {
			err = fmt.Errorf("manifest not found: %s", manifest)
			break
		}
		list = append(list, manifest)
	}
	return list, err
}
