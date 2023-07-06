package repotools

type GetProjectsFilter byte

const (
	Enabled GetProjectsFilter = iota
	Disabled
	All
)

func (o *GetProjectsFilter) String() string {
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
