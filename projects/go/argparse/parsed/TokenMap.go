package parsed

import (
	"github.com/sam-caldwell/monorepo/v2/projects/go/argparse/parsed/token"
)

// Namespace - a map of Token objects
type Namespace struct {
	data map[string]token.Token
}
