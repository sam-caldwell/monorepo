package volatileSecrets

import "github.com/sam-caldwell/monorepo/go/random"

var (
	paddingA         uint
	paddingB         uint
	paddingC         uint
	bitfieldPrefixSz uint
)

func init() {
	//initialize our password bitfield padding states.
	paddingA = random.Uint(secretSzLen, 2*secretSzLen)
	paddingB = random.Uint(secretSzLen, 2*paddingA)
	paddingC = random.Uint(paddingB, 2*paddingB)
}
