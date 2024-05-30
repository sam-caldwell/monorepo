package main

import (
	"fmt"
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
		header.Print()
		if err.Error() == "EOF" {
			ansi.Green().Println("OK").Reset()
		} else {
			ansi.Red().Printf("Error:%s", err).LF().Reset().Fatal(exit.GeneralError)
		}
	}

	header.Print()
	return
	for _, record := range records {
		fmt.Printf(""+
			"Record Type: %d,\n"+
			"Flags: %d,\n"+
			"Data Size: %d\n",
			record.RecordType, record.Flags, record.DataSize)
		fmt.Printf("Data: %x\n\n", record.Data)
	}
}
