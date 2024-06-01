package main

import (
	"github.com/sam-caldwell/monorepo/go/ansi"
	"github.com/sam-caldwell/monorepo/go/crypto/gpg/keybox"
	"github.com/sam-caldwell/monorepo/go/exit"
	"os"
)

/*
   See https://github.com/gpg/gnupg/blob/master/kbx/keybox-blob.c
*/

func main() {
	var (
		err     error
		records []keybox.KBXRecord
		header  keybox.KBXHeader
	)
	kbxFilePath := func() string { // Path to your .kbx file
		if len(os.Args) < 2 {
			ansi.Red().
				Println("Usage: readGnuPgKbxFile <pathFileName.kbx>").
				Fatal(exit.MissingArg)
		}
		p := os.Args[1]
		ansi.Green().Printf("Loading keys from %s\n", p).LF().Reset()
		return p
	}()

	if header, records, err = keybox.ParseKBXFile(kbxFilePath); err != nil {
		if err.Error() == "EOF" {
			ansi.Green().Println("OK").Reset()
		} else {
			header.Print()
			ansi.Red().Printf("Error:%s", err).LF().Reset().Fatal(exit.GeneralError)
		}
	}
	header.Print()
	for i, record := range records {
		recordTypeStr := "unknown"
		if record.Header.RecordType == 02 {
			recordTypeStr = "OpenPGP"
		}
		ansi.Cyan().LF().
			Printf("Record #%02d\n", i).
			Printf("Length: %d (0x%02x)\n", record.Header.DataSize, record.Header.DataSize).
			Printf("Record Type: 0x%02x (%s),\n", record.Header.RecordType, recordTypeStr).
			Printf("Flags: 0x%02x,\n", record.Header.BlobFlags).
			Printf("Version: %d\n", record.Header.Version).
			Printf("DataOffset: %d (0x%02x)\n", record.Header.DataOffset, record.Header.DataOffset).
			Printf("DataLength: %d (0x%02x)\n", record.Header.DataLength, record.Header.DataLength).
			Printf("KeyCount: %d (0x%02x)\n", record.Header.KeyCount, record.Header.KeyCount).
			Printf("KeyInfoLength: %d (0x%02x)\n", record.Header.KeyInfoLength, record.Header.KeyInfoLength).
			Reset()
		for i, key := range record.Keys {
			ansi.Cyan().
				Printf("\tKeyFpr[%d]: 0x%02x\n", i, key.KeyFpr).
				Printf("\tKidOffset[%d]: 0x%02x (%d)\n", i, key.KidOffset, key.KidOffset).
				Printf("\tKeyFlags[%d]: 0x%02x (%d)\n", i, key.KeyFlags, key.KeyFlags).
				Printf("\tReserved[%d]: 0x%02x\n", i, key.Reserved1).
				Printf("\tSzOfSerialNumber[%d]: 0x%02x (%d)\n", i, key.SzOfSerialNumber, key.SzOfSerialNumber).
				Reset()
		}

		//fmt.Printf("Data: %x\n\n", record.Body.Data)
	}
}
