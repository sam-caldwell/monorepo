package JiraActions

import (
	"flag"
	Atlassian "github.com/sam-caldwell/monorepo/go/atlassian"
)

func getExpand() *string {
	return flag.String(
		"expand",
		Atlassian.DefaultExpand,
		"Use expand to include additional information about issues in the response.")
}
