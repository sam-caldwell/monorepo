package repolinter

func (linter *Linters) Initialize() {
	linter.table = map[string]LinterFunction{
		extAsmAmd64:   AssemblyAmd64,
		extAsmArm64:   AssemblyArm64,
		extCSource:    C,
		extCHeader:    C,
		extCppSource:  Cpp,
		extCppHeader:  Cpp,
		extGolang:     Golang,
		extNode:       NodeJs,
		extPython:     Python,
		extMarkdown:   Markdown,
		extShell:      Shell,
		extTypeScript: TypeScript,
		extYaml:       Yaml,
	}
}
