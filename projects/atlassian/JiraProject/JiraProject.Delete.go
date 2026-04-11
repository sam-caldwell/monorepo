package JiraProject

import (
	Atlassian "github.com/sam-caldwell/monorepo/projects/atlassian"
	"net/http"
)

// Delete - Delete the given project
func (jira *Project) Delete(domain *Atlassian.Domain, project string) (*http.Request, error) {
	return nil, nil
}
