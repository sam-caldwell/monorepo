package packageManager

// String - Return string form of InstallAs type
func (o *InstallAs) String() string {
	switch *o {
	case Pkg:
		return "Package"
	case Shell:
		return "Shell"
	case GoGet:
		return "GoGet"
	case GoInstall:
		return "GoInstall"
	default:
		return "Unsupported/Unknown"
	}
}
