package args

import (
	cs "github.com/knightsc/gapstone"
	"github.com/sam-caldwell/monorepo/go/ansi"
	"github.com/sam-caldwell/monorepo/go/exit"
	"github.com/sam-caldwell/monorepo/go/misc/words"
)

func (data *Data) SetArch(arch *string) {
	switch *arch {

	case words.Amd32:
		data.arch = cs.CS_ARCH_X86
		data.mode = cs.CS_MODE_32

	case words.Arm32:
		data.arch = cs.CS_ARCH_ARM
		data.mode = cs.CS_MODE_32

	case words.Amd64:
		data.arch = cs.CS_ARCH_X86
		data.mode = cs.CS_MODE_64

	case words.Arm64:
		data.arch = cs.CS_ARCH_ARM
		data.mode = cs.CS_MODE_ARM

	default:
		ansi.Red().
			Printf("invalid or unsupported architecture").
			LF().
			Reset().
			Fatal(exit.InvalidInput)
	}
}
