package keyvault

import (
	"encoding/base64"
	"fmt"
	"github.com/sam-caldwell/monorepo/go/fs/file"
)

// Save - Save the KeyVault object to disk.
//
//	 Compress, then Base64 encode the ciphertext
//	 Write the result to memory
//
//		(c) 2023 Sam Caldwell.  MIT License
func (kv *KeyVault) Save() (err error) {
	var vault file.File
	if err = vault.Open(kv.fileName, file.FlagWriteOnly|file.FlagTruncate, file.OwnerWrite|file.OwnerRead); err != nil {
		return err
	}
	content := fmt.Sprintf(headerPattern, kv.Version.Major, kv.Version.Minor) + "\n"

	if kv.Payload.cipher == nil {
		panic("nil cipher")
	}

	for i, c := range base64.StdEncoding.EncodeToString(kv.Payload.cipher) {
		content += string(c)
		if (i+1)%textWidth == 0 {
			content += "\n"
		}
	}

	content += fmt.Sprintf(footerPattern, kv.Version.Major, kv.Version.Minor) + "\n"

	if _, err = vault.WriteBytes([]byte(content)); err != nil {
		return err
	}
	return err
}
