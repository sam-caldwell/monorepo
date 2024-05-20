package keyvalue

import (
	"fmt"
	"github.com/sam-caldwell/monorepo/go/exit/errors"
	"github.com/sam-caldwell/monorepo/go/keyvalue/pair"
	"os"
)

// WriteFile - Write the KeyValue store to disk.
//
//	     No file will be created if KeyValue is not initialized.  An error will be returned.
//
//
//		(c) 2023 Sam Caldwell.  MIT License
func (kv *KeyValue[KeyType, ValueType]) WriteFile(fileName string, columnDelimiter, lineEnding rune) (err error) {
	var (
		file             *os.File
		columnDelimBytes []byte
		lineEndingBytes  []byte
	)

	if kv.data == nil {
		return fmt.Errorf(errors.UninitializedValue)
	}

	if columnDelimBytes, err = pair.KeyToBytes(columnDelimiter); err != nil {
		return err
	}

	if lineEndingBytes, err = pair.KeyToBytes(lineEnding); err != nil {
		return err
	}

	if file, err = os.Create(fileName); err != nil {
		return fmt.Errorf("error creating file(%v)", err)
	}

	kv.lock.RLock()
	defer func() {
		kv.lock.RUnlock()
		if file != nil {
			_ = file.Close()
		}
	}()

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

		if _, err = file.Write(columnDelimBytes); err != nil {
			return fmt.Errorf("error writing column delimiter to file: %w", err)
		}

		if _, err = file.Write(valueBytes); err != nil {
			return fmt.Errorf("error writing value to file: %w", err)
		}

		if _, err = file.Write(lineEndingBytes); err != nil {
			return fmt.Errorf("error writing line ending to file: %w", err)
		}
	}
	return err
}
