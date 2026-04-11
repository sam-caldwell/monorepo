package keyvalue

import (
	"github.com/sam-caldwell/monorepo/projects/misc/words"
)

func Interceptor[KeyType comparable, ValueType any](sourceFunc func() (KeyValue[KeyType, ValueType], error)) (string, error) {
	kv, err := sourceFunc()
	return kv.ToString(words.Colon, words.NewLine), err
}
