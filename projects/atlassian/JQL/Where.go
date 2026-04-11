package JQL

// Where - Jql Query where clause
type Where struct {
	Field      *Field      `json:"field,omitempty"`
	Clauses    []Clause    `json:"clauses,omitempty"`
	Operator   string      `json:"operator,omitempty"`
	Predicates []Predicate `json:"predicates,omitempty"`
}
