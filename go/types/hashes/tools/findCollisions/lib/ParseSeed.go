package findCollision

import (
	"encoding/hex"
	"log"
)

// ParseSeed - parse raw hex-encoded string for byte values representing a seed number
func ParseSeed(raw *string) ([]byte, error) {
	if *raw == "" {
		log.Fatal("a seed must be specified")
	}
	decoded, err := hex.DecodeString(*raw)
	if err != nil {
		log.Fatalf("decoding error: %v", err)
	}
	length := len(decoded)
	reversed := make([]byte, length)
	for i := 0; i < length; i++ {
		reversed[length-i-1] = decoded[i]
	}
	return reversed, err
}
