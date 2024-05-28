package LogEvent

// MessageKeyValue - Nested log message data (scalar)
//
//	(c) 2023 Sam Caldwell.  MIT License
type MessageKeyValue[ValueType comparable] struct {
	Field string    `json:"field"`
	Value ValueType `json:"value"`
}
