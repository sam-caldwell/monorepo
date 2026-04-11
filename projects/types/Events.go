package types

import (
	"fmt"
	"time"
)

// Event - Struct describes a single event captured by the system
type Event struct {
	// This object defines an event when it is pushed into the eventQueue
	// serialized for transmission.
	EventTime time.Time  //Event occurred when?
	SendTime  time.Time  //Event sent to server when?
	Resource  ResourceId //numeric resource Id
	Field     FieldId    //numeric field Id
	Value     []byte     //serialized value
}

func (evt *Event) ToString() string {
	return fmt.Sprintf("{"+
		" EventTime:'%s',"+
		" SendTime:'%s',"+
		" Resource:%d,"+
		" Field:%d,"+
		" Value:%v"+
		"}",
		evt.EventTime, evt.SendTime,
		evt.Resource, evt.Field,
		evt.Value)
}
