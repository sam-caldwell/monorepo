package JiraIssue

import "github.com/sam-caldwell/monorepo/go/ansi"

//Update -  Update the identified issue as defined by the internal Issue struct state
func (jira Issue) Update() (any, error) {
    ansi.Blue().Print("Update issue").LF().Reset()
    return nil, nil
}
