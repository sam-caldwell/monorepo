package keyvalue

import (
	"bytes"
	"encoding/binary"
	"fmt"
)

// ValueToBytes converts a value of ValueType to a []byte.
func ValueToBytes[ValueType any](value ValueType) ([]byte, error) {
	var buf bytes.Buffer
	switch v := any(value).(type) {
	case int, int8, int16, int32, int64:
		err := binary.Write(&buf, binary.LittleEndian, v)
		if err != nil {
			return nil, fmt.Errorf("error converting value to bytes: %w", err)
		}
	case uint, uint8, uint16, uint32, uint64:
		err := binary.Write(&buf, binary.LittleEndian, v)
		if err != nil {
			return nil, fmt.Errorf("error converting value to bytes: %w", err)
		}
	case float32, float64:
		err := binary.Write(&buf, binary.LittleEndian, v)
		if err != nil {
			return nil, fmt.Errorf("error converting value to bytes: %w", err)
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
		return nil, fmt.Errorf("unsupported value type: %T", value)
	}
	return buf.Bytes(), nil
}
