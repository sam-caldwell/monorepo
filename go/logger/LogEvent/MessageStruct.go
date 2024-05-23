package LogEvent

// MessageStruct - Nested log message data (struct)
type MessageStruct[ValueType MessageValue] struct {
	Field string    `json:"field"`
	Value ValueType `json:"value"`
}
