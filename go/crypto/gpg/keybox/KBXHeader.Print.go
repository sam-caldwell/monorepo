package keybox

import (
	"github.com/sam-caldwell/monorepo/go/ansi"
)

// Print - Print the header to stdout.
func (header *KBXHeader) Print() {
	ansi.LF().
		Green().Printf("KBX Header:\n").
		Green().Printf("\t     Magic: ").
		Blue().Printf("0x%02x (%s)\n", header.Magic, header.Magic).
		Green().Printf("\tRecordSize: ").
		Blue().Printf("0x%02x (%d)\n", header.RecordSize, header.RecordSize).
		Green().Printf("\t   Version: ").
		Blue().Printf("%d.%d\n", header.VersionMajor, header.VersionMajor).
		Green().Printf("\t     Flags: ").
		Blue().Printf("0x%04x (%d)\n", header.Flags, header.Flags).
		Green().Printf("\t     RecordType: ").
		Blue().Printf("0x%04x (%d)\n", header.RecordType, header.RecordType).
		Green().Printf("\t CreatedAt: ").
		Blue().Printf("0x%02x (%d)\n", header.CreatedAt, header.CreatedAt).
		Green().Printf("\t LastMaint: ").
		Blue().Printf("0x%02x (%d)\n", header.LastMaint, header.LastMaint).
		Green().Printf("\t  Reserved: ").
		Blue().Printf("0x%02x (%d)\n", header.Reserved, header.Reserved).
		LF().
		Reset()
}
