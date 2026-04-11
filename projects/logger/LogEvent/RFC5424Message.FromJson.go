package LogEvent

import "encoding/json"

// FromJson - unmarshal a JSON byte slice to RFC5424Message
//
//	(c) 2023 Sam Caldwell.  MIT License
func (e *RFC5424Message) FromJson(data []byte) error {
	return json.Unmarshal(data, &e)
}
