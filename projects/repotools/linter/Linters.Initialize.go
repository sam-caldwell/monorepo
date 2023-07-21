package repolinter

import (
	"github.com/sam-caldwell/go/v2/projects/ansi"
	"github.com/sam-caldwell/go/v2/projects/exit"
	"github.com/sam-caldwell/go/v2/projects/misc/words"
	systemrecon "github.com/sam-caldwell/go/v2/projects/systemrecon/opsys"
)

func (linter *Linters) Initialize() {
	linter.runners = map[string]LinterFunction{
		extAsmAmd64:   AssemblyAmd64,
		extAsmArm64:   AssemblyArm64,
		extCSource:    C,
		extCHeader:    C,
		extCppSource:  Cpp,
		extCppHeader:  Cpp,
		extGolang:     Golang,
		extJson:       Json,
		extNode:       NodeJs,
		extPython:     Python,
		extMarkdown:   Markdown,
		extShell:      Shell,
		extTypeScript: TypeScript,
		extYaml:       Yaml,
		extYml:        Yaml,
	}
	linter.commands = map[string]string{
		extAsmAmd64:   cmdAsmAmd64,
		extAsmArm64:   cmdAsmArm64,
		extCSource:    cmdC,
		extCHeader:    cmdC,
		extCppSource:  cmdCpp,
		extCppHeader:  cmdCpp,
		extGolang:     cmdGolang,
		extJson:       cmdJson,
		extNode:       cmdNodeJs,
		extPython:     cmdPython,
		extMarkdown:   cmdMarkdown,
		extShell:      cmdShell,
		extTypeScript: cmdTypeScript,
		extYaml:       cmdYaml,
		extYml:        cmdYaml,
	}
	for extension, command := range linter.commands {
		if command == "" {
			continue //we can't really validate a non-existent linter
		}
		if _, response := systemrecon.HasExecutable(command); response == words.Yes {
			continue
		} else {
			ansi.
				Red().
				Printf("Error: %s is not installed to process %s", command, extension).
				LF().
				Reset().
				Fatal(exit.GeneralError)
		}
	}
}
