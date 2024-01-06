package pkicore

import (
	"encoding/pem"
	"fmt"
)

// ParsePEMBlock parses a PEM-encoded block and returns the decoded bytes.
func ParsePEMBlock(pemData []byte, expectedType string) ([]byte, error) {
	block, _ := pem.Decode(pemData)
	if block == nil || block.Type != expectedType {
		return nil, fmt.Errorf("failed to decode PEM block with type %s", expectedType)
	}
	return block.Bytes, nil
}
