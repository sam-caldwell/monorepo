package JiraIssue

import "github.com/sam-caldwell/monorepo/go/ansi"

// Create - create issue defined by the internal Issue struct state
func (jira Issue) Create() (any, error) {
    ansi.Blue().Print("create issue").LF().Reset()
    return nil, nil
}
