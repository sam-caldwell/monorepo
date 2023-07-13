package monorepo

const (
	// CommandDirectory - this is the directory for projects that produce executables and other binary artifacts.
	CommandDirectory = "cmd"

	// ProjectsDirectory - this is the directory for reusable code projects.
	ProjectsDirectory = "projects"

	// NumberOfPartsInCommandPath - the number of path parts in a command path (cmd/<project>/<program>/Manifest.yaml)
	NumberOfPartsInCommandPath = 4

	// NumberOfPartsInProjectPath  - the number of path-parts in a projects path (projects/<project>/Manifest.yaml)
	NumberOfPartsInProjectPath = 3
)
