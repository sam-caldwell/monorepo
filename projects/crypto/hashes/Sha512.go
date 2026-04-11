package hashes

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

// Sha512Bytes - Return Sha512 hash of an input byte slice
func Sha512Bytes(input []byte) []byte {
	hash := sha512.New()
	hash.Write(input)
	return hash.Sum(nil)
}
