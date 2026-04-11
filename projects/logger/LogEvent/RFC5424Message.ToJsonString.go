package LogEvent

// ToJsonString - marshals RFC5424Message to a JSON string
//
//	(c) 2023 Sam Caldwell.  MIT License
func (e *RFC5424Message) ToJsonString() string {
	return string(e.ToJson())
}
