package Disassember

import (
	cs "github.com/knightsc/gapstone"
)

// recursiveDisassembly performs recursive disassembly
func recursiveDisassembly(binary []byte, arch int, mode int, entryPoint uint64) (result []cs.Instruction, err error) {

	var engine cs.Engine

	engine, err = cs.New(arch, mode)
	if err != nil {
		return nil, err
	}

	defer func() { _ = engine.Close() }()

	address := entryPoint

	for len(binary) > 0 {
		instrs, err := engine.Disasm(binary, address, 0)
		if err != nil {
			return nil, err
		}
		if len(instrs) == 0 {
			break
		}
		result = append(result, instrs[0])
		instrSize := instrs[0].Size
		binary = binary[instrSize:]
		address += uint64(instrSize)
	}
	return result, nil
}
