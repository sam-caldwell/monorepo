package keyvalue

import (
	"fmt"
	"github.com/sam-caldwell/monorepo/go/exit/errors"
)

// Walk - For each record in KeyValue, execute the provided walker function to process the record.
//
//	(c) 2023 Sam Caldwell.  MIT License
func (kv *KeyValue[KeyType, ValueType]) Walk(fn func(key KeyType, value ValueType) error) error {
	if kv.data == nil {
		return fmt.Errorf(errors.UninitializedValue)
	}
	kv.lock.Lock()
	defer kv.lock.Unlock()
	for key, value := range kv.data {
		if err := fn(key, value); err != nil {
			return err
		}
	}
	return nil
}
