package parsed

import (
	"github.com/sam-caldwell/go-monorepo/v2/projects/argparse/parsed/token"
)

// Namespace - a map of Token objects
type Namespace struct {
	data map[string]token.Token
}
