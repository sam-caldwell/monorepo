package volatileSecrets

import "github.com/sam-caldwell/monorepo/go/types/hashes"

const (
	secretSzLen uint = 32 / 8              //number bytes used to encode the secret size
	keySz       uint = hashes.Sha512Length // 64 bytes
)
