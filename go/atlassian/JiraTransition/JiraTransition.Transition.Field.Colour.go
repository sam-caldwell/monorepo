package JiraTransition

// Colour - Represent Transition.Field.Colour to give details about the transition
type Colour struct {
	AllowedValues   []string `json:"allowedValues"`
	DefaultValue    string   `json:"defaultValue"`
	HasDefaultValue bool     `json:"hasDefaultValue"`
	Key             string   `json:"key"`
	Name            string   `json:"name"`
	Operations      []string `json:"operations"`
	Required        bool     `json:"required"`
	Schema          Schema   `json:"schema"`
}
