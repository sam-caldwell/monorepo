package JiraProject

import (
	Atlassian "github.com/sam-caldwell/monorepo/go/atlassian"
	AtlassianTypes "github.com/sam-caldwell/monorepo/go/atlassian/types"
	"net/http"
)

// List - List all projects (paginated)
func (jira *Project) List(domain *Atlassian.Domain, jql *AtlassianTypes.JqlQuery) (*http.Request, error) {
	return nil, nil
}
