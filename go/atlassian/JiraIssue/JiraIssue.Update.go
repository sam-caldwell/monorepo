package JiraIssue

import (
    "github.com/sam-caldwell/monorepo/go/ansi"
    Atlassian "github.com/sam-caldwell/monorepo/go/atlassian"
    "net/http"
)

//Update -  Update the identified issue as defined by the internal Issue struct state
func (jira Issue) Update(domain *Atlassian.Domain, web *http.Client) (*http.Request, error) {
    ansi.Blue().Print("Update issue").LF().Reset()
    return nil, nil
}
