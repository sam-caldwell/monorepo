package types

type SystemEvent struct {
	// This is the raw event which will be serialized into []byte
}

func (evt *SystemEvent) Serialize() []byte {
	return nil
}
func (evt *SystemEvent) DeSerialize([]byte) {}
