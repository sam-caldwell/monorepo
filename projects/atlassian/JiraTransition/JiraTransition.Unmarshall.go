package JiraTransition

import "encoding/json"

// Unmarshall - unmarshall the JSON data into a Jira transition
func (jira *JiraTransition) Unmarshall(data []byte) error {
	err := json.Unmarshal(data, &jira.transitions)
	if err != nil {
		return err
	}
	return nil
}
