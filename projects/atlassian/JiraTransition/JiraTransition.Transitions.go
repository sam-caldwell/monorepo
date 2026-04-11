package JiraTransition

// Transitions - Represent a collection of transitions
type Transitions struct {
	Expand      string       `json:"expand,omitempty"`
	Transitions []Transition `json:"transitions"`
}
