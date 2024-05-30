package keybox

import (
	"encoding/binary"
	"github.com/sam-caldwell/monorepo/go/ansi"
	"os"
)

// ReadKBXRecord - Read a KBX key record
//
//	(c) 2023 Sam Caldwell.  MIT License
func ReadKBXRecord(file *os.File) (record KBXRecord, err error) {

	ansi.Green().Println("Reading RecordType").Reset()
	if err = binary.Read(file, binary.BigEndian, &record.RecordType); err != nil {
		return KBXRecord{}, err
	}

	ansi.Green().Println("Reading Flags").Reset()
	if err = binary.Read(file, binary.BigEndian, &record.Flags); err != nil {
		return KBXRecord{}, err
	}

	ansi.Green().Println("Reading DataSize").Reset()
	if err = binary.Read(file, binary.BigEndian, &record.DataSize); err != nil {
		return KBXRecord{}, err
	}

	ansi.Green().Println("Reading Data").Reset()
	record.Data = make([]byte, record.DataSize)
	if _, err = file.Read(record.Data); err != nil {
		return KBXRecord{}, err
	}
	return record, nil
}
