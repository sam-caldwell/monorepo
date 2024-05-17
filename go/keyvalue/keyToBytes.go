package keyvalue

import (
	"bytes"
	"encoding/binary"
	"fmt"
)

// KeyToBytes converts a key of KeyType to a []byte.
func KeyToBytes[KeyType any](key KeyType) ([]byte, error) {
	var buf bytes.Buffer
	switch v := any(key).(type) {
	case int, int8, int16, int32, int64:
		err := binary.Write(&buf, binary.LittleEndian, v)
		if err != nil {
			return nil, fmt.Errorf("error converting key to bytes: %w", err)
		}
	case uint, uint8, uint16, uint32, uint64:
		err := binary.Write(&buf, binary.LittleEndian, v)
		if err != nil {
			return nil, fmt.Errorf("error converting key to bytes: %w", err)
		}
	case float32, float64:
		err := binary.Write(&buf, binary.LittleEndian, v)
		if err != nil {
			return nil, fmt.Errorf("error converting key to bytes: %w", err)
		}
	case string:
		buf.WriteString(v)
	case []byte:
		buf.Write(v)
	case bool:
		var b byte
		if v {
			b = 1
		} else {
			b = 0
		}
		buf.WriteByte(b)
	default:
		return nil, fmt.Errorf("unsupported key type: %T", key)
	}
	return buf.Bytes(), nil
}
