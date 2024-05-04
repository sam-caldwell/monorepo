package JiraIssue

import (
	"encoding/json"
	"fmt"
	"github.com/sam-caldwell/monorepo/go/fs/file"
	"os"
)

// Load - Load the Jira Issue JSON file
func (jira Issue) Load(fileName string) error {
	if !file.Exists(fileName) {
		return fmt.Errorf("descriptor file does not exist")
	}
	content, err := os.ReadFile(fileName)
	if err != nil {
		return err
	}
	err = json.Unmarshal(content, &jira)
	if err != nil {
		return err
	}
	//
	//ansi.Blue().Printf("load descriptor: %s", fileName).LF().Reset()
	return nil
}
