package JiraIssue

import (
	"bytes"
	"fmt"
	Atlassian "github.com/sam-caldwell/monorepo/go/atlassian"
	"net/http"
)

// Update -  Update the identified issue as defined by the internal Issue struct state
func (jira *Issue) Update(domain *Atlassian.Domain, issueIdOrKey string) (*http.Request, error) {
	const path = "/rest/api/3/issue/%s" // %s is {issueIdOrKey}

	return http.NewRequest(
		http.MethodPut,
		Atlassian.JiraUrlFactory(
			Atlassian.JiraUrlPattern,
			domain.Get(),
			fmt.Sprintf(path, issueIdOrKey)),
		bytes.NewBuffer(jira.Marshall()))
}
