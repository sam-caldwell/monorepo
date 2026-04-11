package JiraProject

import "encoding/json"

// Unmarshall - unmarshall the JSON data into a Jira Project
func (jira *Project) Unmarshall(data []byte) error {
	err := json.Unmarshal(data, jira)
	if err != nil {
		return err
	}
	return nil
}
