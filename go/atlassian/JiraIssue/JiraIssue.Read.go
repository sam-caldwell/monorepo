package JiraIssue

import (
	"bytes"
	"fmt"
	Atlassian "github.com/sam-caldwell/monorepo/go/atlassian"
	"net/http"
)

// Read - read the specified issue
func (jira *Issue) Read(domain *Atlassian.Domain, issueIdOrKey string) (*http.Request, error) {
	const path = "/rest/api/3/issue/%s" // %s is {issueIdOrKey}

	return http.NewRequest(
		http.MethodGet,
		Atlassian.JiraUrlFactory(
			Atlassian.JiraUrlPattern,
			domain.Get(),
			fmt.Sprintf(path, issueIdOrKey)),
		bytes.NewBuffer(jira.Marshall()))

}
