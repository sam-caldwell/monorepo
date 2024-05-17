package keyvalue

import (
	"fmt"
	"github.com/sam-caldwell/monorepo/go/keyvalue/pair"
	"os"
)

// WriteFile - Write the KeyValue store to disk.
func (kv *KeyValue[KeyType, ValueType]) WriteFile(fileName string, columnDelimiter, lineEnding rune) (err error) {
	kv.lock.RLock()
	defer kv.lock.RUnlock()

	file, err := os.Create(fileName)
	if err != nil {
		return fmt.Errorf("error creating file: %w", err)
	}
	defer func() { _ = file.Close() }()

	for key, value := range kv.data {
		var keyBytes, valueBytes []byte

		if keyBytes, err = pair.KeyToBytes(key); err != nil {
			return fmt.Errorf("error converting key to bytes: %w", err)
		}

		if valueBytes, err = pair.ValueToBytes(value); err != nil {
			return fmt.Errorf("error converting value to bytes: %w", err)
		}

		if _, err = file.Write(keyBytes); err != nil {
			return fmt.Errorf("error writing key to file: %w", err)
		}

		if _, err = file.Write([]byte{byte(columnDelimiter)}); err != nil {
			return fmt.Errorf("error writing column delimiter to file: %w", err)
		}

		if _, err = file.Write(valueBytes); err != nil {
			return fmt.Errorf("error writing value to file: %w", err)
		}

		if _, err = file.Write([]byte{byte(lineEnding)}); err != nil {
			return fmt.Errorf("error writing line ending to file: %w", err)
		}
	}
	return err
}
