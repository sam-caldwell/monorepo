package threatQL

import "encoding/json"

// Bytes - Return the serialized bytes for a given query struct
func (q *Query) Bytes() (result []byte, err error) {
	return json.Marshal(*q)
}
