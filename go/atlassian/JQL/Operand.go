package JQL

// Operand - JQL Clause Operand
type Operand struct {
	EncodedName    string     `json:"encodedName,omitempty"`
	Name           string     `json:"name,omitempty"`
	EncodedValue   string     `json:"encodedValue,omitempty"`
	Value          string     `json:"value,omitempty"`
	EncodedOperand string     `json:"encodedOperand,omitempty"`
	Values         []Value    `json:"values,omitempty"`
	Arguments      []string   `json:"arguments,omitempty"`
	Function       string     `json:"function,omitempty"`
	Property       []Property `json:"property,omitempty"`
}
