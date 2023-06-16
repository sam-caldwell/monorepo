package parsed

import (
	"github.com/sam-caldwell/argparse/v2/argparse/parsed/token"
)

// Namespace - a map of Token objects
type Namespace struct {
	data map[string]token.Token
}
