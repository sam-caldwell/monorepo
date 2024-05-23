package LogEvent

// MessageKeyValue - Nested log message data (scalar)
type MessageKeyValue[ValueType comparable] struct {
	Field string    `json:"field"`
	Value ValueType `json:"value"`
}
