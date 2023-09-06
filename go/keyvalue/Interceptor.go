package keyvalue

import (
	"github.com/sam-caldwell/monorepo/go/misc/words"
)

func Interceptor(sourceFunc func() (KeyValue, error)) (string, error) {
	kv, err := sourceFunc()
	return kv.ToString(words.Colon, words.NewLine), err
}
