package file

import (
	"os"
)

const (

	/*
	   File Flags
	*/

	//FlagAppend - can be appended
	FlagAppend = os.O_APPEND

	//FlagCreate - can be created if not exists
	FlagCreate = os.O_CREATE

	//FlagExclusive - opened exclusively
	FlagExclusive = os.O_EXCL

	//FlagReadOnly - open for readonly access
	FlagReadOnly = os.O_RDONLY

	//FlagReadWrite - open file for read/write random access
	FlagReadWrite = os.O_RDWR

	//FlagSync - open file for synchronous access
	FlagSync = os.O_SYNC

	//FlagTruncate - open file and truncate the content
	FlagTruncate = os.O_TRUNC

	//FlagWriteOnly - open file for write-only access
	FlagWriteOnly = os.O_WRONLY
)
