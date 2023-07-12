package main

import (
	"flag"
	"github.com/sam-caldwell/go/v2/projects/ansi"
	"github.com/sam-caldwell/go/v2/projects/exit"
	"github.com/sam-caldwell/go/v2/projects/fs/directory"
	"github.com/sam-caldwell/go/v2/projects/fs/file"
	"github.com/sam-caldwell/go/v2/projects/repotools/manifest" // Replace with your actual package import path
	reposkeleton "github.com/sam-caldwell/go/v2/projects/repotools/skeleton"
	systemrecon "github.com/sam-caldwell/go/v2/projects/systemrecon/opsys"
	"path/filepath"
	"runtime"
	"strings"
)

func main() {

	const (
		ManifestFile         = "MANIFEST.yaml"
		projectRootDirectory = "projects"
	)

	whoami, _ := systemrecon.GetCurrentUserName()
	author := flag.String("author", whoami, "Project owner/author")
	project := flag.String("project", "", "The new project name")
	language := flag.String("language", "go", "The programming language")

	manifestOnly := flag.Bool("manifest-only", false, "Only add a manifest file")
	flag.Parse()

	if strings.TrimSpace(*author) == "" {
		*author = whoami
	}

	projectName := strings.TrimSpace(*project)
	if projectName == "" {
		ansi.Red().Printf("Empty project name not allowed").
			LF().
			Fatal(exit.GeneralError)
	}
	projectDirectory := filepath.Clean(filepath.Join(projectRootDirectory, projectName))

	ansi.Blue().
		Tab().Printf("  author: %s", *author).LF().
		Tab().Printf(" project: %s", projectName).LF().
		Tab().Printf("language: %s", *language).LF().
		Reset()

	if *manifestOnly {
		if !directory.Exists(projectDirectory) {
			ansi.Red().
				Println("manifest-only option cannot be used if the project does not exist.").
				LF().
				Fatal(exit.GeneralError)
		}
	} else {
		err := reposkeleton.CreateSkeleton().
			Language(*language).
			Directory(projectDirectory).
			ReadMeFile().
			Commit().
			Error()
		if err != nil {
			ansi.Red().
				Printf("Error creating skeleton (%s): %s", projectDirectory, err).
				LF().
				Fatal(exit.GeneralError)
		}
	}
	manifestFileName := filepath.Join(projectDirectory, ManifestFile)
	if file.Exists(manifestFileName) {
		ansi.Red().
			Printf("manifest file exists: %s", manifestFileName).
			LF().
			Fatal(exit.GeneralError)
	}

	err := projectmanifest.
		CreateManifest(manifestFileName).
		SetName(projectName).
		SetAuthor(*author).
		SetLanguage("go").
		EnableLint().
		DisableBuild().
		DisablePack().
		DisableScan().
		DisableSign().
		AddPlatform(runtime.GOOS, runtime.GOARCH).
		AddPlatform("linux", "amd64").
		AddPlatform("windows", "amd64").
		AddDependency("ansi/color").
		AddDependency("exit").
		AddDependency("misc").
		WriteFile().
		Error()

	if err != nil {
		ansi.Red().Printf("Error: %s", err).LF().Fatal(exit.GeneralError)
	}

	ansi.
		Green().
		Printf("Created %s", manifestFileName).LF().
		Println("(This manifest file is a standard template.  Edits are required)").
		LF().
		Reset()
}
