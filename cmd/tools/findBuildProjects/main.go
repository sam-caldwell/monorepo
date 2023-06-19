package main

import (
	"fmt"
	"github.com/sam-caldwell/go/v2/projects/fs/file"
	"path/filepath"
)

func main() {
	const (
		rootDirectory      = "cmd"
		disabledSignalFile = "build.disabled"
	)

	files, err := filepath.Glob(fmt.Sprintf("%s/*/*/main.go", rootDirectory))
	if err != nil {
		fmt.Printf("Error: %s\n", err)
		return
	}
	for _, thisFile := range files {
		var directory string

		directory = filepath.Dir(thisFile)
		project := filepath.Base(filepath.Dir(filepath.Dir(thisFile)))

		if file.Exists(filepath.Join(rootDirectory, project, disabledSignalFile)) {
			continue // the project is disabled
		}
		directory = filepath.Dir(thisFile)
		program := filepath.Base(directory)

		if file.Exists(filepath.Join(rootDirectory, project, program, disabledSignalFile)) {
			continue // the program is disabled
		}
		fmt.Printf("%s/%s/%s\n", rootDirectory, project, program)

	}
}
