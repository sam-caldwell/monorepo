package Disassember

import (
	cs "github.com/knightsc/gapstone"
	"github.com/sam-caldwell/monorepo/go/ansi"
)

// linearDisassembly performs Linear disassembly
func linearDisassembly(debug bool, binary []byte, arch int, mode int, entryPoint uint64) (result []cs.Instruction, err error) {

	var engine cs.Engine

	if debug {
		ansi.Cyan().Println("using linear disassembly").Reset()
	}

	engine, err = cs.New(arch, mode)
	if err != nil {
		return nil, err
	}

	defer func() { _ = engine.Close() }()

	return engine.Disasm(binary, entryPoint, 0)
}
