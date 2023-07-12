package repotools

type ProjectType byte

const (
	CommandProject ProjectType = 0
	PackageProject ProjectType = 1
)

func (o ProjectType) String() string {
	switch o {
	case CommandProject:
		return "Command"
	case PackageProject:
		return "Package"
	default:
		panic("Invalid Project Type")
	}
}
