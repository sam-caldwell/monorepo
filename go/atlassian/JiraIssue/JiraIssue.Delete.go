package JiraIssue

import (
    "github.com/sam-caldwell/monorepo/go/ansi"
    Atlassian "github.com/sam-caldwell/monorepo/go/atlassian"
    "net/http"
)

// Delete the specified issue
func (jira Issue) Delete(domain *Atlassian.Domain, web *http.Client) (*http.Request, error) {
    ansi.Blue().Print("Delete issue").LF().Reset()
    return nil, nil
}
