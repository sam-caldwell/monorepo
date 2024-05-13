package JQL

// Query - Represents a JQL Query
type Query struct {
	Query     string    `json:"query,omitempty"`
	Structure Structure `json:"structure,omitempty"`
	Errors    []string  `json:"errors,omitempty"`
}
