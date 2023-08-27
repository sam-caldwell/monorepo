package packageManager

// DependencyDescriptor - A descriptor representing a single dependency
type DependencyDescriptor struct {
	Name           string    // Logical name used for logging
	Installer      InstallAs //Installer to use
	Detail         string    // The parameters passed to the installer
	SkipIfDetected bool      //Do not install if package detected. (noop if Shell/GoGet)
}
