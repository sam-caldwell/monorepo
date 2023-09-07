package args

import (
	"fmt"
	"github.com/sam-caldwell/monorepo/go/exit"
)

func failIfMissingExpectedValue(t string) {
	exit.OnCondition(
		true,
		exit.InvalidInput,
		fmt.Sprintf("Expected value but got %s", t),
		checkUsage)
}
