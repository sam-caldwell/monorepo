package executables

// Elf32 - 32-bit executable binary format header structure
//
//	(c) 2023 Sam Caldwell.  MIT License
type Elf32 struct {
}

//typedef struct
//{
//unsigned char	e_ident[EI_NIDENT];	/* Magic number and other info */
//Elf32_Half	e_type;			/* Object file type */
//Elf32_Half	e_machine;		/* Architecture */
//Elf32_Word	e_version;		/* Object file version */
//Elf32_Addr	e_entry;		/* Entry point virtual address */
//Elf32_Off	e_phoff;		/* Program header table file offset */
//Elf32_Off	e_shoff;		/* Section header table file offset */
//Elf32_Word	e_flags;		/* Processor-specific flags */
//Elf32_Half	e_ehsize;		/* ELF header size in bytes */
//Elf32_Half	e_phentsize;		/* Program header table entry size */
//Elf32_Half	e_phnum;		/* Program header table entry count */
//Elf32_Half	e_shentsize;		/* Section header table entry size */
//Elf32_Half	e_shnum;		/* Section header table entry count */
//Elf32_Half	e_shstrndx;		/* Section header string table index */
//} Elf32_Ehdr;
