package JiraProject

import (
	"encoding/json"
	"os"
)

// Load - Load the Jira Project JSON file
func (jira *Project) Load(fileName string) error {
	content, err := os.ReadFile(fileName)
	if err != nil {
		return err
	}
	if err = json.Unmarshal(content, &jira); err != nil {
		return err
	}
	return nil
}
