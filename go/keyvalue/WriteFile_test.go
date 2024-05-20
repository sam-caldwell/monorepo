package keyvalue

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/sam-caldwell/monorepo/go/exit/errors"
	"github.com/sam-caldwell/monorepo/go/fs/file"
	"os"
	"testing"
)

func TestKeyValue_WriteFile(t *testing.T) {
	t.Run("test with uninitialized file", func(t *testing.T) {
		var kv KeyValue[string, string]
		if kv.data != nil {
			t.Fatalf("expected nil data")
		}
		fileName := fmt.Sprintf("/tmp/%s.txt", uuid.New().String())
		err := kv.WriteFile(fileName, ':', '\n')
		if err == nil {
			t.Fatalf("expected error, got none")
		}
		if err.Error() != errors.UninitializedValue {
			t.Fatalf("expected uninitialized value error, got %v", err)
		}
		if file.Exists(fileName) {
			t.Fatalf("file exists but should not have been created (%s)", fileName)
		}
	})
	t.Run("test with initialized data set", func(t *testing.T) {
		var kv KeyValue[string, string]
		fileName := fmt.Sprintf("/tmp/%s.txt", uuid.New().String())
		t.Cleanup(func() {
			if err := os.Remove(fileName); err != nil {
				t.Fatal(err)
			}
		})

		t.Run("setup test data", func(t *testing.T) {
			kv.data = map[string]string{
				"key1": "value1",
				"key2": "value2",
				"key3": "value3",
				"key4": "value4",
			}
		})

		t.Run("create test file", func(t *testing.T) {
			err := kv.WriteFile(fileName, ':', '\n')
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
		})

		t.Run("verify test file exists", func(t *testing.T) {
			if !file.Exists(fileName) {
				t.Fatalf("file not created (%s)", fileName)
			}
		})

		t.Run("read test file", func(t *testing.T) {
			content, err := os.ReadFile(fileName)
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			expected := []byte{
				//each line is 12 bytes long, when encoded.
				//00    01    02    03    04    05    06    07    08    09    0A    0B    0C    0D    0E    0F    10    11
				//k    e     y     1     :     nil   nil   nil   v     a     l     u     e     1     \n    nil   nil   nil
				0x6b, 0x65, 0x79, 0x31, 0x3a, 0x00, 0x00, 0x00, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x31, 0x0a, 0x00, 0x00, 0x00,
				//k    e     y     2     :     nil   nil   nil   v     a     l     u     e     2     \n    nil   nil   nil
				0x6b, 0x65, 0x79, 0x32, 0x3a, 0x00, 0x00, 0x00, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x32, 0x0a, 0x00, 0x00, 0x00,
				//k    e     y     3     :     nil   nil   nil   v     a     l     u     e     3     \n    nil   nil   nil
				0x6b, 0x65, 0x79, 0x33, 0x3a, 0x00, 0x00, 0x00, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x33, 0x0a, 0x00, 0x00, 0x00,
				//k    e     y     4     :     nil   nil   nil   v     a     l     u     e     4     \n    nil   nil   nil
				0x6b, 0x65, 0x79, 0x34, 0x3a, 0x00, 0x00, 0x00, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x34, 0x0a, 0x00, 0x00, 0x00,
			}
			if string(content) != string(expected) {
				t.Fatalf("content mismatch:\n"+
					"actual:\n"+
					"\n'%02x'\n"+
					"\n'%s'\n"+
					"expected:\n"+
					"\n'%02x'\n"+
					"\n'%s'\n", []byte(content), string(content), []byte(expected), string(expected))
			}
		})
	})
}
