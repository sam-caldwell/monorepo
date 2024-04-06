package types

import (
	"time"
)

type Event struct {
	// This object defines an event when it is pushed into the eventQueue
	// serialized for transmission.
	Time     time.Time
	Resource ResourceId
	Field    FieldId
	Value    []byte
}
