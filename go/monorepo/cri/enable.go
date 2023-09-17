package monorepoCri

import (
	"github.com/sam-caldwell/monorepo/go/misc/words"
	criCommon "github.com/sam-caldwell/monorepo/go/monorepo/common"
)

// Enable - Change CRI config to enabled
func Enable(name string) error {
	return criCommon.UpdateField(
		words.ContainerRuntimes,
		words.Name,
		name,
		words.Enabled,
		true)
}
