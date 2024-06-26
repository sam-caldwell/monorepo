package JQL

// SearchQuery - Represent a Jira query request for the search API endpoint
type SearchQuery struct {
	Expand       []string `json:"expand"`
	Fields       []string `json:"fields"`
	FieldsByKeys bool     `json:"fieldsByKeys"`
	Jql          string   `json:"jql"`
	MaxResults   uint     `json:"maxResults"`
	StartAt      uint     `json:"startAt"`
}
