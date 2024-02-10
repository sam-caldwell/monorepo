package Disassember

import (
	"fmt"
	cs "github.com/knightsc/gapstone"
	"github.com/sam-caldwell/monorepo/go/fs/file"
	"os"
)

// Disassemble - Given a source and output file, disassemble the source binary and write the assembly instructions.
func Disassemble(sourceFile, outFile *os.File, method DisassemblerMethod, arch, mode int) (err error) {
	//
	// arch, mode should be something like cs.CS_ARCH_X86, cs.CS_MODE_64
	//
	const entryPoint uint64 = 0x1000

	var buffer []byte
	var engine cs.Engine
	var instructions []cs.Instruction
	var disassembleFunc func([]byte, int, int, uint64) ([]cs.Instruction, error)

	if err = func() error {
		// Read the entire source file
		buffer = make([]byte, file.GetFileSize(sourceFile))
		_, err = sourceFile.Read(buffer)
		defer func() { _ = sourceFile.Close() }()
		return err
	}(); err != nil {
		return err
	}

	// Initialize Gapstone disassembler
	if engine, err = cs.New(arch, mode); err != nil {
		return fmt.Errorf("error initializing disassembler: %v", err)
	}
	defer func() { _ = engine.Close() }()

	switch method {
	case linear:
		disassembleFunc = linearDisassembly
	case recursive:
		disassembleFunc = recursiveDisassembly
	default:
		return fmt.Errorf("invalid disassembly method. Please choose 'linear' or 'recursive'")
	}

	if instructions, err = disassembleFunc(buffer, arch, mode, entryPoint); err != nil {
		return fmt.Errorf("disassembly error: %v", err)
	}

	for _, instr := range instructions {
		_, err := fmt.Fprintf(outFile, "%s\t%s\n", instr.Mnemonic, instr.OpStr)
		if err != nil {
			return fmt.Errorf("error writing to output file: %v", err)
		}
	}

	return nil
}
