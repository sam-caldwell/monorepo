package crypto

import (
	"crypto/sha512"
	"encoding/hex"
)

// Sha512 - Return Sha512 hash of an input string
func Sha512(input string) string {
	hash := sha512.New()
	hash.Write([]byte(input))
	hashInBytes := hash.Sum(nil)
	hashString := hex.EncodeToString(hashInBytes)
	return hashString
}
