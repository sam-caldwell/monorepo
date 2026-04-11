package JiraConfig

// Property - Represents a single configuration property
type Property struct {
	DefaultValue string `json:"defaultValue"`
	Desc         string `json:"desc"`
	Id           string `json:"id"`
	Key          string `json:"key"`
	Name         string `json:"name"`
	Type         string `json:"type"`
	Value        string `json:"value"`
}
