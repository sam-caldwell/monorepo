package reposkeleton

import (
	"fmt"
	"github.com/sam-caldwell/go/v2/projects/exit/errors"
	"github.com/sam-caldwell/go/v2/projects/fs/directory"
	"os"
	"path/filepath"
	"strings"
)

type RepoSkeleton struct {
	err           error
	language      string
	directory     string
	readMeContent []string
}

func CreateSkeleton() *RepoSkeleton {
	skel := RepoSkeleton{}
	return &skel
}

func (skeleton *RepoSkeleton) Directory(projectDirectory string) *RepoSkeleton {
	if skeleton.err == nil {
		if directory.Exists(projectDirectory) {
			skeleton.err = fmt.Errorf("directory exists:%s", projectDirectory)
		} else {
			skeleton.directory = projectDirectory
		}
	}
	return skeleton
}

func (skeleton *RepoSkeleton) Language(language string) *RepoSkeleton {
	if skeleton.err == nil {
		switch thisLanguage := strings.TrimSpace(strings.ToLower(language)); thisLanguage {
		case "amd64Asm", "arm64Asm", "c", "cpp", "go", "node", "python", "rust", "typescript":
			skeleton.language = thisLanguage
		default:
			skeleton.err = fmt.Errorf(errors.UnsupportedLanguage)
		}
	}
	return skeleton
}

func (skeleton *RepoSkeleton) ReadMeFile(content ...string) *RepoSkeleton {
	if len(content) == 0 {
		content = []string{
			"ProjectName Goes Here",
			"=====================",
			"",
			"## Description",
			"Describe the project briefly",
		}
	}
	if skeleton.err == nil {
		skeleton.readMeContent = append(skeleton.readMeContent, content...)
	}

	return skeleton
}

func (skeleton *RepoSkeleton) Commit() (r *RepoSkeleton) {
	if skeleton.err == nil {
		if err := os.MkdirAll(skeleton.directory, 0755); err != nil {
			skeleton.err = err
			return
		}
		//Readme file
		content := []byte(strings.Join(skeleton.readMeContent, "\n"))
		err := os.WriteFile(filepath.Join(skeleton.directory, "README.md"), content, 0644)
		if err != nil {
			skeleton.err = err
		}
	}
	return skeleton
}

func (skeleton *RepoSkeleton) Error() error {
	return skeleton.err
}
