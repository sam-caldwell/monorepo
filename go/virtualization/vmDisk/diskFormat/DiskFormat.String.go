package diskFormat

import "github.com/sam-caldwell/monorepo/go/misc/words"

func (format *DiskFormat) String() string {
	switch *format {
	case Undefined:
		return words.Undefined
	case Ext4:
		return "ext4"
	case Swap:
		return "swap"
	default:
		panic("unknown or unrecognized disk format value")
	}
}
