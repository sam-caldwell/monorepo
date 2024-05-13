package JQL

// Field - Jql Query OrderBy Field
type Field struct {
	Direction string `json:"direction,omitempty"`
	Field     struct {
		EncodedName string `json:"encodedName,omitempty"`
		Name        string `json:"name,omitempty"`
	} `json:"field,omitempty"`
}
