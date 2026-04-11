package types

type FileSystemEvent struct {
	// This is the raw event which will be serialized into []byte
}

func (evt *FileSystemEvent) Serialize() []byte {
	return nil
}
func (evt *FileSystemEvent) DeSerialize([]byte) {}
