package JiraIssue

import (
    "github.com/sam-caldwell/monorepo/go/ansi"
    Atlassian "github.com/sam-caldwell/monorepo/go/atlassian"
    "net/http"
)

// Read - read the specified issue
func (jira Issue) Read(domain *Atlassian.Domain) (*http.Request, error) {
    ansi.Blue().Print("Read issue").LF().Reset()
    return nil, nil
}
