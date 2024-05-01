package JiraIssue

import "github.com/sam-caldwell/monorepo/go/ansi"

// List - list issues
func (jira Issue) List() (any, error) {
    ansi.Blue().Print("List issue").LF().Reset()
    return nil, nil
}
