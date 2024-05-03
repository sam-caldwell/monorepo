package JiraIssue

import (
    "bytes"
    "fmt"
    Atlassian "github.com/sam-caldwell/monorepo/go/atlassian"
    "net/http"
)

// Create - create issue defined by the internal Issue struct state
func (jira Issue) Create(domain *Atlassian.Domain) (*http.Request, error) {

    return http.NewRequest(
        http.MethodPost,
        fmt.Sprintf(Atlassian.JiraUrlPattern, domain.Get(), Atlassian.JiraApiIssue),
        bytes.NewBuffer(jira.Marshall()))

}
