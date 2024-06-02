package bitreader

import "fmt"

// mockByteReader implements io.ByteReader for testing purposes.
type mockByteReader struct {
	data []byte
	pos  int
}

func (m *mockByteReader) ReadByte() (byte, error) {
	if m.pos >= len(m.data) {
		return 0, fmt.Errorf("end of data")
	}
	b := m.data[m.pos]
	m.pos++
	return b, nil
}
