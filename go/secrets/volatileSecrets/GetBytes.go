package volatileSecrets

import (
	"encoding/binary"
	"github.com/sam-caldwell/monorepo/go/crypto"
)

// GetBytes - Retrieve the secret in memory (yeah...this will put it in cleartext in memory for a time).
// Tip: keep the returned variable local and remove it as soon as possible.
func (p *Password) GetBytes() []byte {
	//
	// parse the uint32 secret size value
	//
	szStart := uint(0)
	szStop := szStart + secretSzLen
	secretSz := uint(binary.BigEndian.Uint32(p.data[szStart:szStop]))
	//log.Printf(""+
	//	"    szStart: %04d "+
	//	"     szStop: %04d [ v:%d]", szStart, szStop, secretSz)
	//
	// get the start and stop boundary for the key/passphrase
	//
	keyStart := paddingA + szStop
	keyStop := keyStart + keySz
	//log.Printf(""+
	//	"   keyStart: %04d "+
	//	"    keyStop: %04d [sz:%d]", keyStart, keyStop, keyStop-keyStart)
	//
	// get the start and stop boundary for the secret
	//
	secretStart := keyStop + paddingB
	secretStop := secretStart + secretSz
	//log.Printf(""+
	//	"secretStart: %04d "+
	//	" secretStop: %04d [sz:%d]", secretStart, secretStop, secretStop-secretStart)
	//
	// decrypt the secret using the passphrase
	//
	//log.Printf("inputs:\n"+
	//	"\t    key: (%04d) %02x\n"+
	//	"\t cipher: %02x\n",
	//	len(p.data[keyStart:keyStop]),
	//	p.data[keyStart:keyStop],
	//	p.data[secretStart:secretStop])
	s, err := crypto.DecryptWithPassphrase(p.data[secretStart:secretStop], p.data[keyStart:keyStop])
	if err != nil {
		panic(err)
	}
	return s
}
