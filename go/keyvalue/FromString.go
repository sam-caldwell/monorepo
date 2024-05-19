package keyvalue

// FromString - Given a reference to a string, parse by lines and key-value columns, storing internally.
//
//	(c) 2023 Sam Caldwell.  MIT License
func (kv *KeyValue[KeyType, ValueType]) FromString(data *string, lineEnding, columnDelimiter string) error {
	buffer := []byte(*data)
	return kv.FromBytes(&buffer, lineEnding, columnDelimiter)
}
