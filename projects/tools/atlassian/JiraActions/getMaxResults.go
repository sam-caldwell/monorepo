package JiraActions

import (
	"flag"
	Atlassian "github.com/sam-caldwell/monorepo/projects/atlassian"
)

func getMaxResults() *uint {
	return flag.Uint("maxResults", Atlassian.MaxResults, "Max results per page")
}
