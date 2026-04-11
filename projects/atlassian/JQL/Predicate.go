package JQL

// Predicate - Jql Where clause predicate
type Predicate struct {
	Operand  Operand `json:"operand,omitempty"`
	Operator string  `json:"operator,omitempty"`
}
