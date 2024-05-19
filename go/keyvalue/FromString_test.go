package keyvalue

import (
	"github.com/sam-caldwell/monorepo/go/misc"
	"testing"
)

func TestFromString(t *testing.T) {
	t.Run("Happy path", func(t *testing.T) {
		kv := KeyValue[string, string]{}
		data := "key1:value1\nkey2:value2\nkey3:value3"
		lineEnding := "\n"
		columnDelimiter := ":"
		if err := kv.FromString(&data, lineEnding, columnDelimiter); err != nil {
			t.Fatal(err)
		}
		expectedData := KeyValue[string, string]{
			data: map[string]string{
				"key1": "value1",
				"key2": "value2",
				"key3": "value3",
			},
		}
		if !misc.EqualMap(&kv.data, &expectedData.data) {
			t.Errorf("Parsed data does not match expected data")
		}
	})

}
