package repolinter

import (
	"github.com/sam-caldwell/go/v2/projects/ansi"
	"github.com/sam-caldwell/go/v2/projects/exit"
	"github.com/sam-caldwell/go/v2/projects/exit/errors"
)

func (linter *Linters) Initialize() {
	linter.table = map[string]LinterConfig{
		extAsmAmd64: {
			name:     "AMD 64 Assembly",
			runner:   AssemblyAmd64,
			preCheck: checkNotConfigured,
		},
		extAsmArm64: {
			name:     "ARM 64 Assembly",
			runner:   AssemblyArm64,
			preCheck: checkNotConfigured,
		},
		extCSource: {
			name:     "C Sources",
			runner:   C,
			preCheck: checkNotConfigured,
		},
		extCHeader: {
			name:     "C Header files",
			runner:   C,
			preCheck: checkNotConfigured,
		},
		extCppSource: {
			name:     "C++ Sources",
			runner:   Cpp,
			preCheck: checkNotConfigured,
		},
		extCppHeader: {
			name:     "C++ Headers",
			runner:   Cpp,
			preCheck: checkNotConfigured,
		},
		extGolang: {
			name:     "Go",
			runner:   Golang,
			preCheck: checkGoLangLint,
		},
		extJson: {
			name:     "JSON",
			runner:   Json,
			preCheck: checkJsonLint,
		},
		extNode: {
			name:     "Node.JS",
			runner:   NodeJs,
			preCheck: checkEsLint,
		},
		extPython: {
			name:     "Python",
			runner:   Python,
			preCheck: checkPythonLinter,
		},
		extMarkdown: {
			name:     "Markdown files",
			runner:   Markdown,
			preCheck: checkNotConfigured,
		},
		extRust: {
			name:     "Rust",
			runner:   Rust,
			preCheck: checkNotConfigured,
		},
		extShell: {
			name:     "Bash Shell Scripts",
			runner:   Shell,
			preCheck: checkShellCheck,
		},
		extTypeScript: {
			name:     "Typescript",
			runner:   TypeScript,
			preCheck: checkEsLint,
		},
		extYaml: {
			name:     "Yaml files",
			runner:   Yaml,
			preCheck: checkYamlLint,
		},
		extYml: {
			name:     "Yaml files",
			runner:   Yaml,
			preCheck: checkYamlLint,
		},
	}
	for extension, config := range linter.table {
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
