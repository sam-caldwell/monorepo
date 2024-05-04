package JiraIssue

import (
	"bytes"
	"fmt"
	Atlassian "github.com/sam-caldwell/monorepo/go/atlassian"
	"net/http"
)

// Delete the specified issue
func (jira *Issue) Delete(domain *Atlassian.Domain, issueIdOrKey string) (*http.Request, error) {
	const path = "/rest/api/3/issue/%s" // %s is {issueIdOrKey}
	return http.NewRequest(
		http.MethodDelete,
		Atlassian.JiraUrlFactory(
			Atlassian.JiraUrlPattern,
			domain.Get(),
			fmt.Sprintf(path, issueIdOrKey)),
		bytes.NewBuffer(jira.Marshall()))
}
