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

	ansi.Green().Println("Reading record").Reset()
	if err = binary.Read(file, binary.BigEndian, &record.Header); err != nil {
		return KBXRecord{}, err
	}
	ansi.Green().Println("Reading Data").Reset()
	record.Body.Data = make([]byte, record.Header.DataSize)
	if _, err = file.Read(record.Body.Data); err != nil {
		return KBXRecord{}, err
	}
	return record, nil
}
