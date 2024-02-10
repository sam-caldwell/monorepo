package Disassember

import (
	cs "github.com/knightsc/gapstone"
)

// linearDisassembly performs linear disassembly
func linearDisassembly(binary []byte, arch int, mode int, entryPoint uint64) (result []cs.Instruction, err error) {

	var engine cs.Engine

	engine, err = cs.New(arch, mode)
	if err != nil {
		return nil, err
	}

	defer func() { _ = engine.Close() }()

	return engine.Disasm(binary, entryPoint, 0)
}
