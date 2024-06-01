package keybox

import (
	"encoding/binary"
	"github.com/sam-caldwell/monorepo/go/ansi"
	"io"
	"os"
)

// ReadKBXRecord - Read a KBX key record
//
// Note: we assume big-endian for numeric decoding
//
//	     due to RFC2440,RFC4880
//
//		(c) 2023 Sam Caldwell.  MIT License
func ReadKBXRecord(file *os.File) (record KBXRecord, err error) {

	headerStartPos, _ := file.Seek(0, io.SeekCurrent)
	ansi.Green().Printf("Reading record (start: %d)\n", headerStartPos).Reset()
	if err = binary.Read(file, binary.BigEndian, &record.Header); err != nil {
		return KBXRecord{}, err
	}

	//ansi.Green().Printf("Reading record keys\n").Reset()
	//record.Keys = make(map[uint64]KBXRecordKey, record.Header.KeyCount)
	//for i := 0; i < int(record.Header.KeyCount); i++ {
	//	ansi.Green().Printf("\tReading key %d\n", i).Reset()
	//	var KeyData KBXRecordKey
	//	if err = binary.Read(file, binary.BigEndian, &KeyData); err != nil {
	//		return KBXRecord{}, err
	//	}
	//	//The keyId is always 8-bytes from the end of the fingerprint (FPR)
	//	keyId := convert.ByteSliceToUint64WithBigEndian(KeyData.KeyFpr[len(KeyData.KeyFpr)-8:])
	//	record.Keys[keyId] = KeyData
	//}
	//ansi.Green().Printf("key read done\n").LF().Reset()

	//dataStartPos, _ := file.Seek(0, io.SeekCurrent)
	//record.Body.Data = make([]byte, record.Header.DataLength)
	//if _, err = file.Read(record.Body.Data); err != nil {
	//	return KBXRecord{}, err
	//}

	dataEndPos, _ := file.Seek(0, io.SeekCurrent)
	ansi.Red().
		Printf("start: %d end: %d ", dataStartPos, dataEndPos).
		Printf("calcLen: %d ", dataEndPos-dataStartPos).
		Printf("actualLen: %d", record.Header.DataLength).
		LF().Reset()

	return record, nil
}
