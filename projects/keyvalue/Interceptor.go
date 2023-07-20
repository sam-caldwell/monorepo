package keyvalue

import (
	"github.com/sam-caldwell/go/v2/projects/misc/words"
)

func Interceptor(sourceFunc func() (KeyValue, error)) (string, error) {
	kv, err := sourceFunc()
	return kv.ToString(words.Colon, words.NewLine), err
}
