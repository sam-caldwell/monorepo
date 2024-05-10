package JiraTransition

// Summary - Represent the values of Transition.Field.Summary
type Summary struct {
	AllowedValues   []string `json:"allowedValues"`
	DefaultValue    string   `json:"defaultValue"`
	HasDefaultValue bool     `json:"hasDefaultValue"`
	Key             string   `json:"key"`
	Name            string   `json:"name"`
	Operations      []string `json:"operations"`
	Required        bool     `json:"required"`
	Schema          Schema   `json:"schema"`
}
