package executables

// ElfHeaderFieldNames - This type represents the fields from /usr/include/elf.h used w/ GetValue() methods
//
//	(c) 2023 Sam Caldwell.  MIT License
type ElfHeaderFieldNames uint8

const (
	eIdent ElfHeaderFieldNames = iota
	eType
	eMachine
	eVersion
	eEntry
	ePhoff
	eShoff
	eFlags
	eEhSize
	ePhEntSize
	ePhNum
	eShEndSize
	eShNum
	eShStrNdx
)
