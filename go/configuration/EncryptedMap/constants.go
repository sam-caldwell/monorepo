package configuration

import "github.com/sam-caldwell/monorepo/go/fs/file"

const (
	rflags = file.FlagReadOnly
	wflags = file.FlagWriteOnly | file.FlagExclusive

	perms = file.OwnerRead | file.OwnerWrite
)

const (
	errFailedToReadKeyFile = "failed to read key file %s: %w"
	errFailedToReadKeysDir = "failed to read keys directory: %w"
	errNoKeyNotFound       = "no matching key found for identity: %s"
)
