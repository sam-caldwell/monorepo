package volatileSecrets

import (
	"encoding/binary"
	"github.com/sam-caldwell/monorepo/go/crypto"
	"github.com/sam-caldwell/monorepo/go/random"
	"runtime"
)

/*
 * Secret bitfields
 *
 *  [secretSz][paddingA][passphrase][paddingB][secret][passphraseC]
 *
 * Yes, if you get a heap dump you can eventually find the secrets, etc.,
 * and you can parse the bitfield and decrypt them.  But to do so, you
 * will at least have to earn it.  Common tools are gonna be challenged
 * here.
 */

// NewPassword - Create a bitfield over which an encrypted secret will be written
//
//	This is not perfect secrecy, but it is enough to defeat basic hexdump
//	operations. After all, you don't always have to beat the bear, ya just
//	gotta run faster than everyone else running from the bear.
func NewPassword(passphrase []byte, secret []byte) *Password {
	var err error
	var p Password
	//
	// Raw sensitive information
	//
	var cipher, key []byte
	//
	// Make sure we free any memory quickly and safely
	//
	defer func() {
		secureDelete(&key)
		secureDelete(&cipher)
		runtime.GC()
	}()
	//
	//secure the passphrase and secret then free the cleartext.
	//
	key = crypto.Sha512Bytes(passphrase)
	secureDelete(&passphrase)
	//log.Printf("key: %02x", key)
	if cipher, err = crypto.EncryptWithPassphrase(secret, key); err != nil {
		panic(err)
	}
	//log.Printf("secret: %02x", secret)
	secureDelete(&secret)
	//
	//Create a random field over which we will write our secrets...
	//
	secretSz := uint(len(cipher))
	bitfieldSz := int(paddingA + paddingB + paddingC + keySz + secretSz)
	//log.Printf("bitfieldSz: %d", bitfieldSz)
	if p.data, err = random.GenerateRandomBytes(bitfieldSz); err != nil {
		panic(err)
	}
	//log.Printf("bitfieldSz: %d", bitfieldSz)
	//log.Printf("bitfield: %02x", p.data)

	//
	//Overlay the length of the cipher on our bitfield
	//
	startByte := uint(0)
	stopByte := secretSzLen
	binary.BigEndian.PutUint32(p.data[startByte:stopByte], uint32(secretSz))
	//log.Printf("startByte: %d stopByte: %d  [secretSzLen:%d] pA:%d, pB:%d, pC:%d,", startByte, stopByte, secretSzLen, paddingA, paddingB, paddingC)
	//log.Printf("bitfield: %02x", p.data)
	//
	//Overlay the key on the bitfield
	//
	startByte += paddingA + stopByte
	stopByte = startByte + keySz
	copy(p.data[startByte:stopByte], key[:])
	//log.Printf("startByte: %d stopByte: %d  [keySz:%d] pA:%d, pB:%d, pC:%d,", startByte, stopByte, keySz, paddingA, paddingB, paddingC)
	//log.Printf("bitfield: %02x", p.data)
	//
	//Copy the cipher over the random bit field
	//
	startByte = stopByte + paddingB
	stopByte = startByte + uint(len(cipher))
	copy(p.data[startByte:stopByte], cipher[:])
	//log.Printf("startByte: %d stopByte: %d  [secretSz:%d] pA:%d, pB:%d, pC:%d,", startByte, stopByte, secretSz, paddingA, paddingB, paddingC)
	//log.Printf("bitfield: %02x", p.data)
	return &p
}
