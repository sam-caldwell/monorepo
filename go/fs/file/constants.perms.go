package file

import (
	"os"
)

const (
	/*
	   Permission flags.  Use bitwise OR to combine them.
	   This follows the POSIX rwx permission flags (see above).
	*/

	//PermNone - No permissions
	PermNone = 0x00

	//PermRead - Read permissions
	PermRead = 0x04

	//PermWrite - Write permissions
	PermWrite = 0x02

	//PermExecute - Execute permission
	PermExecute = 0x01

	/*
	   Simple Permissions...
	*/

	//OwnerRead - User can read
	OwnerRead = PermRead << AppliesToOwner

	//GroupRead - Group can read
	GroupRead = PermRead << AppliesToGroup

	//EveryoneRead - Everyone can read
	EveryoneRead = PermRead << AppliesToEveryone

	//OwnerWrite - Owner can write
	OwnerWrite = PermWrite << AppliesToOwner

	//GroupWrite - Group can write
	GroupWrite = PermWrite << AppliesToGroup

	//EveryoneWrite - Everyone can write
	EveryoneWrite = PermWrite << AppliesToEveryone

	//OwnerExecute - Owner can execute
	OwnerExecute = PermExecute << AppliesToOwner

	//GroupExecute - Group can execute
	GroupExecute = PermExecute << AppliesToGroup

	//EveryoneExecute - Everyone can execute
	EveryoneExecute = PermExecute << AppliesToEveryone

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
