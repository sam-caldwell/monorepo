package JiraIssue

import (
    "bytes"
    "fmt"
    Atlassian "github.com/sam-caldwell/monorepo/go/atlassian"
    "net/http"
)

// Create - create issue defined by the internal Issue struct state
func (jira Issue) Create(domain *Atlassian.Domain) (*http.Request, error) {

    req, err := http.NewRequest(
        http.MethodPost,
        fmt.Sprintf(Atlassian.JiraUrlPattern, domain.Get(), Atlassian.JiraApiIssue),
        bytes.NewBuffer(jira.Marshall()))

    if err != nil {
        return nil, err
    }

    req.Header.Set("Content-Type", "application/json")

    return req, err

}
