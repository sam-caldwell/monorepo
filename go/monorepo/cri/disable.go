package monorepoCri

import (
	"github.com/sam-caldwell/monorepo/go/misc/words"
	criCommon "github.com/sam-caldwell/monorepo/go/monorepo/common"
)

// Disable - Change CRI config to disabled
func Disable(name string) error {
	return criCommon.UpdateField(
		words.ContainerRuntimes,
		words.Name,
		name,
		words.Enabled,
		false)
}
