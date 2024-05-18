package pair

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"github.com/sam-caldwell/monorepo/go/convert"
	"github.com/sam-caldwell/monorepo/go/exit/errors"
	"reflect"
	"sort"
)

// KeyToBytes converts a key of KeyType to a []byte.
func KeyToBytes[KeyType comparable](key KeyType) ([]byte, error) {
	const conversionError = "error converting key to bytes: %w"
	var buf bytes.Buffer

	if err := writeValue(&buf, key); err != nil {
		return nil, fmt.Errorf(conversionError, err)
	}

	return buf.Bytes(), nil
}

func writeValue(buf *bytes.Buffer, value any) error {
	switch v := value.(type) {
	case int:
		// Decodes any int to 64-bit integer to avoid non-fixed int problem
		if err := binary.Write(buf, binary.LittleEndian, int64(v)); err != nil {
			return fmt.Errorf("error converting key to bytes: %w", err)
		}
	case int8, int16, int32, int64, uint, uint8, uint16, uint32, uint64, float32, float64:
		if err := binary.Write(buf, binary.LittleEndian, v); err != nil {
			return fmt.Errorf("error converting key to bytes: %w", err)
		}
	case string:
		buf.WriteString(v)
	case bool:
		buf.WriteByte(convert.BoolToByte(v))
	case *int:
		if err := binary.Write(buf, binary.LittleEndian, int64(*v)); err != nil {
			return fmt.Errorf("error converting key to bytes: %w", err)
		}
	case *int8:
		if err := binary.Write(buf, binary.LittleEndian, *v); err != nil {
			return fmt.Errorf("error converting key to bytes: %w", err)
		}
	case *int16:
		if err := binary.Write(buf, binary.LittleEndian, *v); err != nil {
			return fmt.Errorf("error converting key to bytes: %w", err)
		}
	case *int32:
		if err := binary.Write(buf, binary.LittleEndian, *v); err != nil {
			return fmt.Errorf("error converting key to bytes: %w", err)
		}
	case *int64:
		if err := binary.Write(buf, binary.LittleEndian, *v); err != nil {
			return fmt.Errorf("error converting key to bytes: %w", err)
		}
	case *uint:
		if err := binary.Write(buf, binary.LittleEndian, *v); err != nil {
			return fmt.Errorf("error converting key to bytes: %w", err)
		}
	case *uint8:
		if err := binary.Write(buf, binary.LittleEndian, *v); err != nil {
			return fmt.Errorf("error converting key to bytes: %w", err)
		}
	case *uint16:
		if err := binary.Write(buf, binary.LittleEndian, *v); err != nil {
			return fmt.Errorf("error converting key to bytes: %w", err)
		}
	case *uint32:
		if err := binary.Write(buf, binary.LittleEndian, *v); err != nil {
			return fmt.Errorf("error converting key to bytes: %w", err)
		}
	case *uint64:
		if err := binary.Write(buf, binary.LittleEndian, *v); err != nil {
			return fmt.Errorf("error converting key to bytes: %w", err)
		}
	case *float32:
		if err := binary.Write(buf, binary.LittleEndian, *v); err != nil {
			return fmt.Errorf("error converting key to bytes: %w", err)
		}
	case *float64:
		if err := binary.Write(buf, binary.LittleEndian, *v); err != nil {
			return fmt.Errorf("error converting key to bytes: %w", err)
		}
	case *string:
		buf.WriteString(*v)
	case *bool:
		buf.WriteByte(convert.BoolToByte(*v))
	case []byte:
		buf.Write(v)
	case *[]byte:
		buf.Write(*v)
	default:
		rv := reflect.ValueOf(value)
		rt := rv.Type()
		if rv.Kind() == reflect.Ptr && rv.Elem().Kind() == reflect.Struct {
			rv = rv.Elem()
			rt = rt.Elem()
		}
		if rv.Kind() == reflect.Struct {
			// Get and sort field names
			fieldNames := make([]string, rv.NumField())
			for i := 0; i < rv.NumField(); i++ {
				fieldNames[i] = rt.Field(i).Name
			}
			sort.Strings(fieldNames)

			// Write fields in sorted order
			for _, fieldName := range fieldNames {
				field := rv.FieldByName(fieldName)
				if err := writeValue(buf, field.Interface()); err != nil {
					return err
				}
			}
			return nil
		}
		return fmt.Errorf(errors.UnsupportedType)
	}
	return nil
}
