package JiraIssue

import "encoding/json"

// Unmarshall - unmarshall the JSON data into a Jira issue
func (jira *Issue) Unmarshall(data []byte) error {
	err := json.Unmarshal(data, jira)
	if err != nil {
		return err
	}
	return nil
}
