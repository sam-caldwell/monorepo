package executables

// Headers - Generic interface for the various executable formats we support
type Headers interface {
	Elf32 | Elf64
}
