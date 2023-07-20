package main

import (
	"fmt"
	"github.com/sam-caldwell/go/v2/projects/exit/errors"
	"github.com/sam-caldwell/go/v2/projects/fs/file"
	"github.com/sam-caldwell/go/v2/projects/lists/SmartList"
	"os"
	"path/filepath"
	"strings"
)

const (
	usage = `
	findProjects [-pretty] [-builds|-tests] <rootDirectory>

			-pretty - an optional argument that prints the results in a pretty format
			-build - find only build-enabled projects
			-test - find only test-enabled projects
			<rootDirectory> - either "cmd" or "projects" to indicate search root.
`

	argCmd            = "cmd"
	argProjects       = "projects"
	buildDisabled     = "build.disabled"
	testDisabled      = "test.disabled"
	searchQuery       = "**/*.go"
	prettyPrintBanner = "current binary projects (%s):\n"
	prettyPrintFormat = "%s/%s/%s\n"
	simplePrintFormat = " - %s/%s/%s\n"
)

//Patterns:
//       <rootDirectory>/<project>...
//  e.g.
//
//  - executables -
//       <rootDirectory>/cmd/<project>/<program>/Build.go
//  - or code modules -
//       <rootDirectory>/cmd/<project[0]>/.../<project[n]>/*.go
//       where the project name is strings.Join(project...)
//
//Project is the grouping of one or more programs or a collection of software

func main() {
	var projects list.SmartList

	rootDirectory, filters, prettyPrint, err := parseArgs()

	exitOnError(err)
	exitOnError(
		filepath.Walk(rootDirectory, func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}
			//We only examine folders
			if info.IsDir() {
				var project string

				if filtered(info.Name(), filters) {
					return nil //Bail! it's a filtered record
				}

				pathParts := strings.Split(path, string(os.PathSeparator))

				depth := len(pathParts)
				if depth < 3 {
					return nil //we're too close to root directory (2 is the project, 3 is the program)
				}

				switch rootDirectory {

				case argCmd:
					if pathParts[0] == argCmd {
						//We need three parts <root0>/<project1>/<program2>
						if depth < 3 {
							return nil
						}
						project = pathParts[1]
						program := pathParts[2]
						if prettyPrint {
							exitOnError(projects.Add(fmt.Sprintf("%s/%s/%s", rootDirectory, project, program)))
						} else {
							exitOnError(projects.Add(fmt.Sprintf("%s/%s/%s", rootDirectory, project, program)))
						}
					}
				case argProjects:
					if pathParts[0] == argProjects {
						exitOnError(projects.Add(filepath.Join(pathParts[1:]...)))
					}
				default:
					return fmt.Errorf(errors.InvalidInput)
				}
			}
			return nil
		}))

	fmt.Println("output:")
	for _, row := range projects.String() {
		fmt.Println(row)
	}
}

// filtered - evaluate thisFile against filenames indicating we should filter thisFile
func filtered(thisFile string, filters []string) (disabled bool) {
	//The list of 'filters' should be a list of filenames we would not want to see
	//in a given directory.
	for _, filterFile := range filters {
		if file.Exists(filepath.Join(thisFile, filterFile)) {
			disabled = true // the project is disabled
		}
	}
	return disabled
}

// printBanner - Print a banner if -pretty was used.
//func printBanner(pretty, filtered bool) {
//	modifier := "unfiltered"
//	if filtered {
//		modifier = "filtered"
//	}
//	if pretty {
//		fmt.Printf(prettyPrintBanner, modifier)
//	}
//}

// printLine - print a line (one project record)
//func printLine(pretty bool, rootDirectory, project, program string) {
//	if pretty {
//		fmt.Printf(prettyPrintFormat, rootDirectory, project, program)
//	} else {
//		fmt.Printf(simplePrintFormat, rootDirectory, project, program)
//	}
//
//}

// exitOnError - on err != nil, exit with error
func exitOnError(err error) {
	if err != nil {
		fmt.Printf("Error: %s\n", err)
		os.Exit(1)
	}
}

// parseArgs - parse this programs command-line arguments and return the state.
func parseArgs() (rootDirectory string, disabledSignalFile []string, prettyPrint bool, err error) {
	if len(os.Args) == 2 {
		if os.Args[1] == "-h" || os.Args[1] == "--help" {
			fmt.Print(usage)
			os.Exit(1)
		}
	}
	if len(os.Args) >= 2 {
		for _, arg := range os.Args[1:] {
			switch strings.ToLower(strings.TrimSpace(arg)) {
			case "--pretty":
				prettyPrint = true
			case "--build":
				disabledSignalFile = append(disabledSignalFile, buildDisabled)
			case "--test":
				disabledSignalFile = append(disabledSignalFile, testDisabled)
			case "cmd":
				fallthrough
			case "projects":
				rootDirectory = arg
			default:
				err = fmt.Errorf("unexpected input: '%s'", arg)
				return
			}
		}
	}
	return rootDirectory, disabledSignalFile, prettyPrint, err
}
