package exit

import (
	"fmt"
	"github.com/sam-caldwell/monorepo/go/misc/words"
	"strings"
)

func ExpectNonEmptyStringP(inp *string) error {
	if inp == nil || strings.TrimSpace(*inp) == words.EmptyString {
		return fmt.Errorf("empty or nil string not allowed")
	}
	return nil
}
