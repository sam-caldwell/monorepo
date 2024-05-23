package LogEvent

import "encoding/json"

// ToJson - marshals RFC5424Message to a JSON byte slice
func (e RFC5424Message) ToJson() ([]byte, error) {
	return json.Marshal(e)
}
