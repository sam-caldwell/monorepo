package JiraIssue

import (
    "bytes"
    "fmt"
    Atlassian "github.com/sam-caldwell/monorepo/go/atlassian"
    "net/http"
)

// Create - create issue defined by the internal Issue struct state
func (jira Issue) Create(domain *Atlassian.Domain, web *http.Client) (*http.Request, error) {

    url := fmt.Sprintf(Atlassian.JiraUrlPattern, domain.Get(), Atlassian.JiraApiIssue)

    return http.NewRequest("POST", url, bytes.NewBuffer(jira.Marshall()))

}
