package AtlassianTypes

// JqlQuery - An object representing a single Jira Query, including the JQL statement and supporting structures.
type JqlQuery struct {
	Expand        []string `json:"expand,omitempty"`
	Fields        []string `json:"fields,omitempty"`
	FieldsByKeys  bool     `json:"fieldsByKeys,omitempty"`
	Jql           string   `json:"jql"` //e.g. "project=Atlassian"
	MaxResults    uint     `json:"maxResults,omitempty"`
	StartAt       uint     `json:"startAt,omitempty"`
	ValidateQuery string   `json:"validateQuery,omitempty"`
}
