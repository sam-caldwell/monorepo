package JiraActions

import (
	"flag"
	Atlassian "github.com/sam-caldwell/monorepo/go/atlassian"
)

func getStartAt() *uint {
	return flag.Uint("startAt", Atlassian.DefaultStartAt, "Offset from the starting record")
}
