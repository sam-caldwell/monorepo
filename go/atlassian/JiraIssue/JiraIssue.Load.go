package JiraIssue

import (
	"encoding/json"
	"fmt"
	"github.com/sam-caldwell/monorepo/go/fs/file"
	"os"
)

// Load - Load the Jira Issue JSON file
func (jira *Issue) Load(fileName *string) error {

	if !file.Existsp(fileName) {
		return fmt.Errorf("descriptor file does not exist")
	}

	content, err := os.ReadFile(*fileName)

	if err != nil {
		return err
	}

	err = json.Unmarshal(content, &jira)

	if err != nil {
		return err
	}

	return nil
}
