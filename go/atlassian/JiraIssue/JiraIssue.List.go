package JiraIssue

import (
    "github.com/sam-caldwell/monorepo/go/ansi"
    Atlassian "github.com/sam-caldwell/monorepo/go/atlassian"
    "net/http"
)

// List - list issues
func (jira Issue) List(domain *Atlassian.Domain) (*http.Request, error) {
    ansi.Blue().Print("List issue").LF().Reset()
    return nil, nil
}
