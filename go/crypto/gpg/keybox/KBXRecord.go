package keybox

// KBXRecord KBX record structure
//
// See https://github.com/gpg/gnupg/blob/master/kbx/keybox-blob.c
//
//	(c) 2023 Sam Caldwell.  MIT License
type KBXRecord struct {
	Header KBXRecordHeader
	Keys   map[uint64]KBXRecordKey
	Body   KBXRecordData
}

type KBXRecordHeader struct {
	DataSize      uint32
	RecordType    uint8
	Version       uint8
	BlobFlags     uint16 //bit 0 = contains secret key/bit 1=ephemeral blob
	DataOffset    uint32 //offset to the OpenPGP key block or X509 DER encoded cert.
	DataLength    uint32 //key/block length
	KeyCount      uint16 //should be (at least one)
	KeyInfoLength uint16 //size of additional key information
}

type KBXRecordKey struct {
	KeyFpr               [20]byte //key fingerprint (MD5 with left-padded 0s)
	KidOffset            uint32   //nth KeyId (KeyId is always 8byte or 0 if not known, eg X509
	KeyFlags             uint16   //special key flag (not yet implemented)
	Reserved1            uint16
	SzOfSerialNumber     uint16
	NumberOfUserIds      uint16
	SzAdditionalUserInfo uint16
	//SerialNumber [??]byte
	//KeyId [8]byte //KeyId

}

type KBXRecordData struct {
	Data []byte
}
