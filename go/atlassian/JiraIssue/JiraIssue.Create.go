package JiraIssue

import (
    "bytes"
    Atlassian "github.com/sam-caldwell/monorepo/go/atlassian"
    "net/http"
)

// Create - create issue defined by the internal Issue struct state
func (jira Issue) Create(domain *Atlassian.Domain) (*http.Request, error) {

    return http.NewRequest(
        http.MethodPost,
        Atlassian.JiraUrlFactory(Atlassian.JiraUrlPattern, domain.Get(), Atlassian.JiraApiIssue),
        bytes.NewBuffer(jira.Marshall()))

}
