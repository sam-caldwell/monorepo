package pair

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"github.com/sam-caldwell/monorepo/go/convert"
)

// ValueToBytes converts a value of ValueType to a []byte.
func ValueToBytes[ValueType any](value ValueType) ([]byte, error) {
	var buf bytes.Buffer
	switch v := any(value).(type) {
	case int:
		if err := binary.Write(&buf, binary.LittleEndian, int64(v)); err != nil {
			return nil, fmt.Errorf("error converting value to bytes: %w", err)
		}
	case uint:
		if err := binary.Write(&buf, binary.LittleEndian, uint64(v)); err != nil {
			return nil, fmt.Errorf("error converting value to bytes: %w", err)
		}
	case int8, int16, int32, int64, uint8, uint16, uint32, uint64, float32, float64:
		if err := binary.Write(&buf, binary.LittleEndian, v); err != nil {
			return nil, fmt.Errorf("error converting value to bytes: %w", err)
		}
	case string:
		buf.WriteString(v)
	case []byte:
		buf.Write(v)
	case bool:
		buf.WriteByte(convert.BoolToByte(v))
	default:
		return nil, fmt.Errorf("unsupported value type: %T", value)
	}
	return buf.Bytes(), nil
}
