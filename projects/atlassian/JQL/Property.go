package JQL

// Property - JQL Query Operand structure property
type Property struct {
	Entity string `json:"entity,omitempty"`
	Key    string `json:"key,omitempty"`
	Path   string `json:"path,omitempty"`
	Type   string `json:"type,omitempty"`
}
