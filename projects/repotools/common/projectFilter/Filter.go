package projectFilter

/*
 * projects/repotools/common/Filter.go
 * (c) 2023 Sam Caldwell.  See License.txt
 *
 * This file defines a filtering mechanism for repotools.ListProjects()
 * which uses a 32-bit unsigned integer as a bit field to represent
 * filter states we will apply when listing the projects in the repo.
 *
 * See README.md
 */

type Filter uint32

/*
 * Note: if you add any new flags, be sure to update FlagList()
 */
const (
	None             Filter = 0x00     //No filter
	Enabled          Filter = 0x01     //    filter flag indicating enabled projects should be shown
	Disabled         Filter = 0x02     //   filter flag indicating disabled projects should be shown
	Windows          Filter = 0x04     //    filter flag indicating Windows support
	Linux            Filter = 0x08     //    filter flag indicating Linux support
	Darwin           Filter = 0x10     //    filter flag indicating macOS support
	Command          Filter = 0x20     //    filter flag indicating projects which compile to executables
	Package          Filter = 0x40     //    filter flag indicating reusable code projects w/ produce no executable
	LanguageGo       Filter = 0x10000  // Indicates project is primarily Golang
	LanguageC        Filter = 0x20000  // Indicates project is primarily C
	LanguageCpp      Filter = 0x40000  // Indicates project is primarily C++
	LanguageAsmAmd64 Filter = 0x80000  // Indicates project is primarily AMD64 (X86_64) Assembly
	LanguageAsmArm64 Filter = 0x100000 // project is primarily Arm64 Assembly
	LanguageRust     Filter = 0x200000 // Indicates project is primarily Rust
	LanguageNodeJs   Filter = 0x400000 // project is primarily NodeJs / Typescript
	LanguagePython   Filter = 0x800000 // Indicates project is primarily Python

)
