package main

import (
	"fmt"
	gpgCrypto "github.com/sam-caldwell/monorepo/go/crypto/gpg"
	"strings"
)

func main() {
	keyDir := "/home/scaldwell/.gnupg/"
	keyIdentity := "0x67B706124DFAABFF"

	publicKeys, err := gpgCrypto.LoadPublicKeys(keyDir, keyIdentity)
	if err != nil {
		panic(err)
	}

	if len(publicKeys) > 0 {
		for _, key := range publicKeys {
			// Print key details
			for _, identity := range key.Identities {
				if strings.Contains(identity.Name, keyIdentity) {
					fingerprint := key.PrimaryKey.Fingerprint
					fmt.Println("Fingerprint:", fingerprint)
					fmt.Println("Identity:", identity.Name)
				}
			}
		}
	} else {
		fmt.Println("No public keys found for the specified identity.")
	}
}
