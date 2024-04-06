package types

type HardwareEvent struct {
	// This is the raw event which will be serialized into []byte
}

func (evt *HardwareEvent) Serialize() []byte {
	return nil
}
func (evt *HardwareEvent) DeSerialize([]byte) {}
