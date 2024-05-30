package main

import (
	"fmt"
	"github.com/sam-caldwell/monorepo/go/ansi"
	gpgCrypto "github.com/sam-caldwell/monorepo/go/crypto/gpg/keybox"
	"github.com/sam-caldwell/monorepo/go/exit"
	"os"
)

/*
   See https://github.com/gpg/gnupg/blob/master/kbx/keybox-blob.c
*/

func main() {
	var err error
	var records []gpgCrypto.KBXRecord
	// Path to your .kbx file
	kbxFilePath := func() string {
		if len(os.Args) < 2 {
			ansi.Red().
				Println("Usage: readGnuPgKbxFile <pathFileName.kbx>").
				Fatal(exit.MissingArg)
		}
		p := os.Args[1]
		ansi.Green().Printf("Loading keys from %s\n", p).LF().Reset()
		return p
	}()

	// Parse the .kbx file
	if records, err = gpgCrypto.ParseKBXFile(kbxFilePath); err != nil {
		ansi.Red().Printf("Error:%s", err).LF().Fatal(exit.GeneralError)
		return
	}

	// Print information about each record
	for _, record := range records {
		fmt.Printf(""+
			"Record Type: %d,\n"+
			"Flags: %d,\n"+
			"Data Size: %d\n",
			record.RecordType, record.Flags, record.DataSize)
		fmt.Printf("Data: %x\n\n", record.Data)
	}
}
