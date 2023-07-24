package repolinter

import (
	"fmt"
	"github.com/sam-caldwell/go/v2/projects/ansi"
	"github.com/sam-caldwell/go/v2/projects/exit"
	"github.com/sam-caldwell/go/v2/projects/exit/errors"
	"os"
)

func (linter *Linters) Initialize(useColor bool) {
	linter.Table = map[string]LinterConfig{
		extAsmAmd64: {
			enabled:        false,
			name:           "AMD 64 Assembly",
			runner:         AssemblyAmd64,
			preCheck:       checkNotConfigured,
			directoryLevel: false,
		},
		extAsmArm64: {
			enabled:        false,
			name:           "ARM 64 Assembly",
			runner:         AssemblyArm64,
			preCheck:       checkNotConfigured,
			directoryLevel: false,
		},
		extCSource: {
			enabled:        false,
			name:           "C Sources",
			runner:         C,
			preCheck:       checkNotConfigured,
			directoryLevel: false,
		},
		extCHeader: {
			enabled:        false,
			name:           "C Header files",
			runner:         C,
			preCheck:       checkNotConfigured,
			directoryLevel: false,
		},
		extCppSource: {
			enabled:        false,
			name:           "C++ Sources",
			runner:         Cpp,
			preCheck:       checkNotConfigured,
			directoryLevel: false,
		},
		extCppHeader: {
			enabled:        false,
			name:           "C++ Headers",
			runner:         Cpp,
			preCheck:       checkNotConfigured,
			directoryLevel: false,
		},
		extGolang: {
			enabled:        true,
			name:           "Go",
			runner:         GolangLinterFactory(),
			preCheck:       checkGoLangLint,
			directoryLevel: true,
		},
		extJson: {
			enabled:        false,
			name:           "JSON",
			runner:         Json,
			preCheck:       checkNotConfigured,
			directoryLevel: false,
		},
		extNode: {
			enabled:        false,
			name:           "Node",
			runner:         NodeJs,
			preCheck:       checkEsLint,
			directoryLevel: false,
		},
		extPython: {
			enabled:        false,
			name:           "Python",
			runner:         Python,
			preCheck:       checkPythonLinter,
			directoryLevel: false,
		},
		extMarkdown: {
			enabled:        false,
			name:           "Markdown",
			runner:         Markdown,
			preCheck:       checkNotConfigured,
			directoryLevel: false,
		},
		extRust: {
			enabled:        false,
			name:           "Rust",
			runner:         Rust,
			preCheck:       checkNotConfigured,
			directoryLevel: false,
		},
		extShell: {
			enabled:        false,
			name:           "Shell",
			runner:         Shell,
			preCheck:       checkShellCheck,
			directoryLevel: false,
		},
		extTypeScript: {
			enabled:        false,
			name:           "Typescript",
			runner:         TypeScript,
			preCheck:       checkEsLint,
			directoryLevel: false,
		},
		extYaml: {
			enabled:        true,
			name:           "Yaml",
			runner:         Yaml,
			preCheck:       checkYamlLint,
			directoryLevel: false,
		},
		extYml: {
			enabled:        true,
			name:           "Yaml",
			runner:         Yaml,
			preCheck:       checkYamlLint,
			directoryLevel: false,
		},
	}
	for extension, config := range linter.Table {
		if err := config.preCheck(); err != nil {
			if err.Error() == errors.NotImplemented {
				message := fmt.Sprintf("Pre-check [SKIP]: %s for %s\n", config.name, extension)
				if useColor {
					ansi.Yellow().Print(message).Reset()
				} else {
					fmt.Print(message)
				}
			} else {
				message := fmt.Sprintf("Pre-check [FAIL]: %s for %s\n", config.name, extension)
				if useColor {
					ansi.Red().Print(message).Reset().Fatal(exit.GeneralError)
				} else {
					fmt.Print(message)
					os.Exit(exit.GeneralError)
				}
			}
		} else {
			message := fmt.Sprintf("Pre-check [PASS]: %s for %s\n", config.name, extension)
			if useColor {
				ansi.Green().Printf(message).Reset()
			} else {
				fmt.Print(message)
				os.Exit(exit.GeneralError)
			}
		}
	}
	if useColor {
		ansi.Green().Println("Linters initialized").Reset()
	}
}
