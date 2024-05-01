package JiraIssue

import "github.com/sam-caldwell/monorepo/go/ansi"

// Read - read the specified issue
func (jira Issue) Read() (any, error) {
    ansi.Blue().Print("Read issue").LF().Reset()
    return nil, nil
}
