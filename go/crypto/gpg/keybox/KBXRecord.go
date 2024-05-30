package keybox

// KBXRecord KBX record structure
//
// See https://github.com/gpg/gnupg/blob/master/kbx/keybox-blob.c
//
//	(c) 2023 Sam Caldwell.  MIT License
type KBXRecord struct {
	DataSize   uint32
	RecordType uint8
	Flags      uint8

	Data []byte
}
