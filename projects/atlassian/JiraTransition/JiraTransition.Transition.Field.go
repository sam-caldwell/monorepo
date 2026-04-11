package JiraTransition

// Field - Represent Transition.Field
type Field struct {
	Summary `json:"summary"`
	Colour  `json:"colour,omitempty"`
}
