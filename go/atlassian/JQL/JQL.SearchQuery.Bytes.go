package JQL

import "encoding/json"

// Bytes - Marshal the internal state into []byte
func (jql *SearchQuery) Bytes() []byte {
	b, err := json.Marshal(jql)
	if err != nil {
		panic(err)
	}
	return b
}
