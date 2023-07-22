package repolinter

import (
	"github.com/sam-caldwell/go/v2/projects/ansi"
	"github.com/sam-caldwell/go/v2/projects/exit"
	"github.com/sam-caldwell/go/v2/projects/exit/errors"
)

func (linter *Linters) Initialize() {
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
			enabled:        true,
			name:           "JSON",
			runner:         Json,
			preCheck:       checkJsonLint,
			directoryLevel: false,
		},
		extNode: {
			enabled:        true,
			name:           "Node",
			runner:         NodeJs,
			preCheck:       checkEsLint,
			directoryLevel: false,
		},
		extPython: {
			enabled:        true,
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
			enabled:        true,
			name:           "Shell",
			runner:         Shell,
			preCheck:       checkShellCheck,
			directoryLevel: false,
		},
		extTypeScript: {
			enabled:        true,
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
				ansi.
					Yellow().
					Printf("Pre-check [SKIP]: %s for %s", config.name, extension).
					LF().
					Reset()
			} else {
				ansi.
					Red().
					Printf("Pre-check [FAIL]: %s for %s", config.name, extension).
					LF().
					Reset().
					Fatal(exit.GeneralError)
			}
		} else {
			ansi.
				Green().
				Printf("Pre-check [PASS]: %s for %s", config.name, extension).
				LF().
				Reset()
		}
	}
}
