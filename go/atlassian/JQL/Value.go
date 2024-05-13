package JQL

// Value - JQL Query Operand Value
type Value struct {
	EncodedValue string `json:"encodedValue,omitempty"`
	Value        string `json:"value,omitempty"`
}
