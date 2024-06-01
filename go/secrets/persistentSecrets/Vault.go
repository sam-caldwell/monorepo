package persistentSecrets

import configuration "github.com/sam-caldwell/monorepo/go/configuration/EncryptedMap"

// Vault - A persistent secret connector based on 1password.
//
//		 This is intended to be an abstraction from 1password to allow us
//	  to pivot to another solution in the future if needed and to avoid any
//	  complex changes needed if 1password changes their sdk.
//
//			(c) 2023 Sam Caldwell.  MIT License
type Vault struct {
	data configuration.EncryptedMap
	//connectHost string
	//token       string
	//vault       uuid.UUID
}
