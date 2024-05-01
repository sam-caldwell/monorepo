package JiraIssue

import "github.com/sam-caldwell/monorepo/go/ansi"

// Delete the specified issue
func (jira Issue) Delete() (any, error) {
    ansi.Blue().Print("Delete issue").LF().Reset()
    return nil, nil
}
