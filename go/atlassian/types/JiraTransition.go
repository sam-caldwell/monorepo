package AtlassianTypes

import Atlassian "github.com/sam-caldwell/monorepo/go/atlassian"

type JiraTransition struct {
	client *Atlassian.Client
	Id     int `json:"id"`
}
