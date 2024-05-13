package JQL

// Clause - Jql Query where clause
type Clause struct {
	Field    Field    `json:"field,omitempty"`
	Operand  Operand  `json:"operand,omitempty"`
	Operator string   `json:"operator,omitempty"`
	Clauses  []Clause `json:"clauses,omitempty"`
}
