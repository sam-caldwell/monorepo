package main

import (
	"flag"
	"github.com/sam-caldwell/go/v2/projects/ansi"
	"github.com/sam-caldwell/go/v2/projects/fs/directory"
	"github.com/sam-caldwell/go/v2/projects/repotools/manifest" // Replace with your actual package import path
	systemrecon "github.com/sam-caldwell/go/v2/projects/systemrecon/opsys"
	"path/filepath"
	"runtime"
	"strings"
)

func main() {
	const projectRootDirectory = "projects"
	whoami, _ := systemrecon.GetCurrentUserName()
	author := flag.String("author", whoami, "Project owner/author")
	project := flag.String("project", "", "The new project name")
	flag.Parse()

	if strings.TrimSpace(*author) == "" {
		*author = whoami
	}

	projectName := strings.TrimSpace(*project)
	if projectName == "" {
		ansi.Red().Printf("Empty project name not allowed").LF().Fatal(1)
	}
	projectDirectory := filepath.Clean(filepath.Join(projectRootDirectory, projectName))
	if directory.Exists(projectDirectory) {
		ansi.Red().Printf("directory exists: %s", projectDirectory).LF().Fatal(1)
	}
	manifestFileName := filepath.Join(projectDirectory, "MANIFEST.yaml")
	// Set the toolchain properties
	err := projectmanifest.
		CreateManifest().
		SetName(projectName).
		SetAuthor(*author).
		SetLanguage("gtxo").
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
		WriteFile(manifestFileName).
		Error()

	if err != nil {
		ansi.Red().Printf("Error: %s", err).LF().Fatal(1)
	}

	ansi.
		Green().
		Printf("Created %s", manifestFileName).LF().
		Println("(This manifest file is a standard template.  Edits are required)").
		LF().
		Reset()
}
