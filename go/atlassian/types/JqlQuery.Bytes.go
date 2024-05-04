package AtlassianTypes

import "encoding/json"

// Bytes - Marshall and return a JSON object as []byte
func (jql *JqlQuery) Bytes() []byte {
	if result, err := json.Marshal(jql); err != nil {
		panic(err)
	} else {
		return result
	}
}
