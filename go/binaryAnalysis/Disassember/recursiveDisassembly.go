package Disassember

import (
	cs "github.com/knightsc/gapstone"
	"github.com/sam-caldwell/monorepo/go/ansi"
)

// recursiveDisassembly performs Recursive disassembly
//
//	(c) 2023 Sam Caldwell.  MIT License
func recursiveDisassembly(debug bool, binary []byte, arch int, mode int, entryPoint uint64) (result []cs.Instruction, err error) {

	var engine cs.Engine

	if debug {
		ansi.Cyan().Println("using recursive disassembly").Reset()
	}

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
