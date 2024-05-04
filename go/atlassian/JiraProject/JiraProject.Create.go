package JiraProject

import (
	Atlassian "github.com/sam-caldwell/monorepo/go/atlassian"
	"net/http"
)

// Create - Create the given project using the internal state
func (jira *Project) Create(domain *Atlassian.Domain) (*http.Request, error) {
	return nil, nil
}
