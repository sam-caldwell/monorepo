package keybox

// KBXHeader KBX header structure
//
// See https://github.com/gpg/gnupg/blob/master/kbx/keybox-blob.c
//
//	 0-31   : Magic Header (32 bits)
//	 32-63  : File Version (32 bits)
//	 64-95  : File Creation Time (32 bits)
//	 96-127 : File Flags (32 bits)
//
//		(c) 2023 Sam Caldwell.  MIT License
type KBXHeader struct {
	RecordSize   uint32
	VersionMajor uint8
	VersionMinor uint8
	RecordType   uint8
	Flags        uint8
	Magic        [8]byte
	CreatedAt    uint32
	LastMaint    uint32
	Reserved     uint64
}
