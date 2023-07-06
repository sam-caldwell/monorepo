package repotools

type projectType byte

const (
	CommandProject projectType = 0
	PackageProject             = 1
)

func (o projectType) String() string {
	switch o {
	case CommandProject:
		return "Command"
	case PackageProject:
		return "Package"
	default:
		panic("Invalid Project Type")
	}
}
