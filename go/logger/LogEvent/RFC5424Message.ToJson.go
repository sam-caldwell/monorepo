package LogEvent

import (
	"encoding/json"
	"github.com/sam-caldwell/monorepo/go/misc/words"
)

// ToJson - marshals RFC5424Message to a JSON byte slice
//
//	(c) 2023 Sam Caldwell.  MIT License
func (e *RFC5424Message) ToJson() []byte {
	b, err := json.MarshalIndent(e, words.EmptyString, words.Tab)
	if err != nil {
		panic(err)
	}
	return b
}
