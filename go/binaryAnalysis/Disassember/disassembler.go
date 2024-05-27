package Disassember

import (
	"fmt"
	cs "github.com/knightsc/gapstone"
	"github.com/sam-caldwell/monorepo/go/ansi"
	"github.com/sam-caldwell/monorepo/go/fs/file"
	"github.com/sam-caldwell/monorepo/go/types"
	"os"
)

// Disassemble - Given a source and output file, disassemble the source binary and write the assembly instructions.
//
//	(c) 2023 Sam Caldwell.  MIT License
func Disassemble(debug bool, sourceFile, outFile *os.File, method types.DisassemblerMethod, arch, mode int) (err error) {
	//
	// arch, mode should be something like cs.CS_ARCH_X86, cs.CS_MODE_64
	//
	const entryPoint uint64 = 0x1000 //WARNING: This is wrong

	var buffer []byte
	var engine cs.Engine
	var instructions []cs.Instruction
	var disassembleFunc func(bool, []byte, int, int, uint64) ([]cs.Instruction, error)

	buffer = make([]byte, file.GetFileSize(sourceFile))

	if debug {
		ansi.Cyan().Printf("Buffer Size: %d", len(buffer)).LF().Reset()
	}

	if _, err = sourceFile.Read(buffer); err != nil {
		return err
	}
	defer func() { _ = sourceFile.Close() }()

	// Initialize Gapstone disassembler
	if engine, err = cs.New(arch, mode); err != nil {
		return fmt.Errorf("error initializing disassembler: %v", err)
	}
	defer func() { _ = engine.Close() }()

	switch method {
	case types.Linear:
		disassembleFunc = linearDisassembly
	case types.Recursive:
		disassembleFunc = recursiveDisassembly
	default:
		return fmt.Errorf("invalid disassembly method. Please choose 'linear' or 'recursive'")
	}

	if instructions, err = disassembleFunc(debug, buffer, arch, mode, entryPoint); err != nil {
		if err.Error() != "OK (CS_ERR_OK)" {
			return fmt.Errorf("disassembly error: %v\n", err)
		}
	}

	if debug {
		ansi.Cyan().Printf("Number of instructions: %d\n", len(instructions)).Reset()
	}

	for _, instr := range instructions {
		line := fmt.Sprintf("%s\t%s\n", instr.Mnemonic, instr.OpStr)
		_, err := fmt.Fprintf(outFile, "%s", line)
		if debug {
			ansi.Cyan().Print(line).Reset()
		}
		if err != nil {
			return fmt.Errorf("error writing to output file: %v", err)
		}
	}

	return nil
}
