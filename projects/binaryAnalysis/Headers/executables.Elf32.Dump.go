package executables

// Dump - Print a string slice of the ELF32 header fields
//
//	(c) 2023 Sam Caldwell.  MIT License
func (elf *Elf32) Dump() []string {
	return []string{
		//fmt.Sprintf("e_ident:     %02x", elf.eIdent),
		//fmt.Sprintf("e_type:      %02x", elf.eType),
		//fmt.Sprintf("e_machine:   %02x", elf.eMachine),
		//fmt.Sprintf("e_version:   %02x", elf.eVersion),
		//fmt.Sprintf("e_entry:     %02x", elf.eEntry),
		//fmt.Sprintf("e_phoff:     %02x", elf.ePhoff),
		//fmt.Sprintf("e_shoff:     %02x", elf.eShoff),
		//fmt.Sprintf("e_flags:     %02x", elf.eFlags),
		//fmt.Sprintf("e_ehsize:    %02x", elf.eEhSize),
		//fmt.Sprintf("e_phentsize: %02x", elf.ePhEntSize),
		//fmt.Sprintf("e_phnum:     %02x", elf.ePhNum),
		//fmt.Sprintf("e_shentsize: %02x", elf.eShEndSize),
		//fmt.Sprintf("e_shnum:     %02x", elf.eShNum),
		//fmt.Sprintf("e_shstrndx:  %02x", elf.eShStrNdx),
	}
}
