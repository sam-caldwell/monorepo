package types

type ProcessEvent struct {
	// This is the raw event which will be serialized into []byte
}

func (evt *ProcessEvent) Serialize() []byte {
	return nil
}
func (evt *ProcessEvent) DeSerialize([]byte) {}
