package JQL

// Structure - Jql Query Structure providing clauses
type Structure struct {
	OrderBy *OrderBy `json:"orderBy,omitempty"`
	Where   *Where   `json:"where,omitempty"`
}
