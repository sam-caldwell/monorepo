package keyvalue

/*
 * keyvalue.FromString
 * (c) 2023 Sam Caldwell.  See LICENSE.txt
 *
 * 	Consume string and parse by lineEnding and columnDelimiter into a keyvalue.Map
 *  internal state.
 *
 *  Note: this is destructive of any existing state
 */

// FromString - Given a reference to a string, parse by lines and key-value columns, storing internally.
func (kv *KeyValue[KeyType, ValueType]) FromString(data *string, lineEnding, columnDelimiter string) error {
	buffer := []byte(*data)
	return kv.FromBytes(&buffer, lineEnding, columnDelimiter)
}
