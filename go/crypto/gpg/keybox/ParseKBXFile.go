package keybox

import (
	"bytes"
	"fmt"
	"github.com/sam-caldwell/monorepo/go/ansi"
	"github.com/sam-caldwell/monorepo/go/fs/file/MagicNumbers"
	"os"
)

// ParseKBXFile - parse the KBX file content
//
//	(c) 2023 Sam Caldwell.  MIT License
func ParseKBXFile(filePath string) (header KBXHeader, records []KBXRecord, err error) {
	var (
		file *os.File
	)
	if file, err = os.Open(filePath); err != nil {
		return header, records, err
	}
	defer func() { _ = file.Close() }()
	if header, err = ReadKBXHeader(file); err != nil {
		return header, records, err
	}
	if !bytes.Equal(header.Magic[:len([]byte(MagicNumbers.GnuPgKbx))], []byte(MagicNumbers.GnuPgKbx)) {
		ansi.Red().
			Printf("   magic:%02x", header.Magic[:len([]byte(MagicNumbers.GnuPgKbx))]).LF().
			Printf("expected:%02x", []byte(MagicNumbers.GnuPgKbx)).LF().
			Reset()
		return header, records, fmt.Errorf("invalid KBX file")
	}
	//for i := 0; i < int(header.NumRecords); i++ { // Read each record
	//	record, err := ReadKBXRecord(file)
	//	if err != nil {
	//		return header, records, err
	//	}
	//	records = append(records, record)
	//}

	return header, records, err
}
