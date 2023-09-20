package cli

import (
	"fmt"
	"github.com/sam-caldwell/monorepo/go/misc/words"
	"go/types"
	"strings"
)

const (
	attrName = "name"
)

const (
	Required = true
	Optional = false
)

type ValueDescriptor struct {
	required bool
	vType    types.Builtin
	value    any
}

func (v *ValueDescriptor) ToString() string {
	return fmt.Sprintf("(value:%v type:%v required:%v) ",
		v.value, v.vType, v.required)
}

type OptionDescriptor struct {
	name     string
	required bool
	vType    types.Builtin
	value    any
}

func (opt *OptionDescriptor) ToString() string {
	return fmt.Sprintf("(name:%v required:%v vType:%v, value: %v)",
		opt.name, opt.required, opt.vType, opt.value)
}

type Command struct {
	descriptor []string
	value      *ValueDescriptor
	options    *OptionDescriptor
}

func (cmd *Command) ToString() string {
	return strings.Join(cmd.descriptor, words.Space) + cmd.value.ToString() + cmd.options.ToString()
}

type ArgumentDescriptor struct {
	//command []Command
	command map[string]Command
}

func (arg *ArgumentDescriptor) ToString() (out string) {
	for _, cmd := range arg.command {
		out = fmt.Sprintf("%s%s\n", out, cmd.ToString())
	}
	return out
}

func NewArgParse() *ArgumentDescriptor {
	return &ArgumentDescriptor{}
}

func (arg *ArgumentDescriptor) Argument(identifiers ...string) *ArgumentDescriptor {
	if len(identifiers) == 0 {
		panic("Empty command identifiers")
	}
	if arg.command == nil {
		arg.command = make(map[string]Command)
	}
	arg.command[strings.Join(identifiers, words.Space)] = Command{}

	return arg
}

func (arg *ArgumentDescriptor) Done() *ArgumentDescriptor {
	return arg
}
