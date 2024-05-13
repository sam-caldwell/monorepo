package JiraActions

import (
	"flag"
	Atlassian "github.com/sam-caldwell/monorepo/go/atlassian"
)

func getFieldsList() *string {
	return flag.String("fields", Atlassian.DefaultFields,
		"A list of fields to return for each issue, use it to retrieve a subset of fields. This parameter accepts a comma-separated list. ")
}
