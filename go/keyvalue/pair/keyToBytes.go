package pair

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"github.com/sam-caldwell/monorepo/go/convert"
)

// KeyToBytes converts a key of KeyType to a []byte.
func KeyToBytes[KeyType comparable](key KeyType) ([]byte, error) {
	var buf bytes.Buffer
	switch v := any(key).(type) {
	case int, int8, int16, int32, int64, uint, uint8, uint16, uint32, uint64, float32, float64:
		err := binary.Write(&buf, binary.LittleEndian, v)
		if err != nil {
			return nil, fmt.Errorf("error converting key to bytes: %w", err)
		}
	case string:
		buf.WriteString(v)
	case bool:
		buf.WriteByte(convert.BoolToByte(v))
	default:
		return nil, fmt.Errorf("unsupported key type: %T", key)
	}
	return buf.Bytes(), nil
}
