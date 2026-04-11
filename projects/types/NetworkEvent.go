package types

type NetworkEvent struct {
	// This is the raw event which will be serialized into []byte
}

func (evt *NetworkEvent) Serialize() []byte {
	return nil
}
func (evt *NetworkEvent) DeSerialize([]byte) {}
