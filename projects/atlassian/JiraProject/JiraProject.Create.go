package JiraProject

import (
	Atlassian "github.com/sam-caldwell/monorepo/projects/atlassian"
	"net/http"
)

// Create - Create the given project using the internal state
func (jira *Project) Create(domain *Atlassian.Domain) (*http.Request, error) {
	return nil, nil
}
