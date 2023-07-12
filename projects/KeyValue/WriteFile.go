package keyvalue

import "fmt"

// WriteFile - Write the key value store to disk.
func (kv *KeyValue) WriteFile(fileName string, columnDelimiter string, lineEnding string) (err error) {
	//ToDo: implement this
	return fmt.Errorf("KeyValue.WriteFile() not implemented.  %s not saved", fileName)
}
