package main

/*
 * create-project
 * (c) 2023 Sam Caldwell.  See LICENSE.txt
 *
 * This program creates a new project within the monorepo
 * as a directory under projects/<project_name> and adds
 * the minimum skeleton (MANIFEST.yaml and README.md).
 */

import (
	"flag"
	ansi2 "github.com/sam-caldwell/go/v2/projects/go/ansi"
	"github.com/sam-caldwell/go/v2/projects/go/exit"
	"github.com/sam-caldwell/go/v2/projects/go/fs/directory"
	"github.com/sam-caldwell/go/v2/projects/go/fs/file"
	"github.com/sam-caldwell/go/v2/projects/go/repotools/manifest"
	"github.com/sam-caldwell/go/v2/projects/go/repotools/skeleton"
	"github.com/sam-caldwell/go/v2/projects/go/systemrecon/opsys"
	"path/filepath"
	"runtime"
	"strings"
)

// main - create-project main function
func main() {

	const (
		projectRootDirectory = "projects"
	)
	exit.IfVersionRequested()

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
		ansi2.Red().Printf("Empty project name not allowed").
			LF().
			Fatal(exit.GeneralError)
	}
	projectDirectory := filepath.Clean(filepath.Join(projectRootDirectory, projectName))

	ansi2.Blue().
		Tab().Printf("  author: %s", *author).LF().
		Tab().Printf(" project: %s", projectName).LF().
		Tab().Printf("language: %s", *language).LF().
		Reset()

	if *manifestOnly {
		if !directory.Exists(projectDirectory) {
			ansi2.Red().
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
			ansi2.Red().
				Printf("Error creating skeleton (%s): %s", projectDirectory, err).
				LF().
				Fatal(exit.GeneralError)
		}
	}
	manifestFileName := filepath.Join(projectDirectory, projectmanifest.ManifestYaml)
	if file.Exists(manifestFileName) {
		ansi2.Red().
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
		ansi2.Red().Printf("Error: %s", err).LF().Fatal(exit.GeneralError)
	}

	ansi2.Green().
		Printf("Created %s", manifestFileName).LF().
		Println("(This manifest file is a standard template.  Edits are required)").
		LF().
		Reset()
}
