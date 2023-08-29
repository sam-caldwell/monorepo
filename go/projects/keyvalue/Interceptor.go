package keyvalue

import (
	words2 "github.com/sam-caldwell/monorepo/go/projects/v2/misc/words"
)

func Interceptor(sourceFunc func() (KeyValue, error)) (string, error) {
	kv, err := sourceFunc()
	return kv.ToString(words2.Colon, words2.NewLine), err
}
