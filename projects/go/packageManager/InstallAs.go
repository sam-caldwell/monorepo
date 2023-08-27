package packageManager

// InstallAs - installation method selector
type InstallAs byte

const (
	// Pkg - Use a system-provided package manager
	Pkg InstallAs = iota

	// Shell - Execute details as shell code
	Shell

	// GoGet - Use 'go get' to fetch a dependency
	GoGet
	// GoInstall - Install dependency with 'go install'
	GoInstall
)
