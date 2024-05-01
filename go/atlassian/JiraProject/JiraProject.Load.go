package JiraProject

import (
    "encoding/json"
    "os"
)

// Load - Load the Jira Project JSON file
func (project JiraProject) Load(fileName string) error {
    content, err := os.ReadFile(fileName)
    if err != nil {
        return err
    }
    if err = json.Unmarshal(content, &project); err != nil {
        return err
    }
    return nil
}
