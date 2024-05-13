package JiraActions

import (
	"flag"
	Atlassian "github.com/sam-caldwell/monorepo/go/atlassian"
)

func getMaxResults() *uint {
	return flag.Uint("maxResults", Atlassian.MaxResults, "Max results per page")
}
