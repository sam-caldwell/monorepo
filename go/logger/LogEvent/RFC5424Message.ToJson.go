package LogEvent

import (
	"encoding/json"
)

// ToJson - marshals RFC5424Message to a JSON byte slice
//
//	(c) 2023 Sam Caldwell.  MIT License
func (e *RFC5424Message) ToJson() []byte {
	b, err := json.Marshal(e)
	if err != nil {
		panic(err)
	}
	return append(b, []byte("\n")...)
}
