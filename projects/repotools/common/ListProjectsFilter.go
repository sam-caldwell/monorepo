package repotools

type ListProjectsFilter byte

const (
	Enabled ListProjectsFilter = iota
	Disabled
	All
)

func (o *ListProjectsFilter) String() string {
	switch *o {
	case Enabled:
		return "Enabled"
	case Disabled:
		return "Disabled"
	case All:
		return "All"
	default:
		return "Invalid"
	}
}
