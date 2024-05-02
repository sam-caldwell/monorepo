package JiraIssue

import "github.com/sam-caldwell/monorepo/go/ansi"

// Load - Load the Jira Issue JSON file
func (jira Issue) Load(fileName string) error {
    ansi.Blue().Printf("load descriptor: %s", fileName).LF().Reset()
    return nil
}
