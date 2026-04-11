package LogEvent

// MessageStruct - Nested log message data (struct)
//
//	(c) 2023 Sam Caldwell.  MIT License
type MessageStruct[ValueType MessageValue] struct {
	Field string    `json:"field"`
	Value ValueType `json:"value"`
}
