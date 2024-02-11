package executables

// Elf64 - 64-bit executable binary format header structure
type Elf64 struct {
	eIdent     [16]byte // Magic number and other info
	eType      uint16   // eType : Object file type
	eMachine   uint16   // e_machine : Architecture
	eVersion   uint32   // e_version : Object file version
	eEntry     uint64   // e_entry : Entry point virtual address
	ePhoff     uint64   // e_phoff : Program header table file offset
	eShoff     uint64   // e_shoff : Section header table file offset
	eFlags     uint32   // e_flags : Processor-specific flags
	eEhSize    uint16   // e_ehsize : ELF header size in bytes
	ePhEntSize uint16   //e_phentsize; Program header table entry size
	ePhNum     uint16   // e_phnum Program header table entry count
	eShEndSize uint16   // e_shentsize : Section header table entry size
	eShNum     uint16   // e_shnum : Section header table entry count
	eShStrNdx  uint16   //	e_shstrndx : Section header string table index
}
