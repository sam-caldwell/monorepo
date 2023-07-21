package repolinter

const (
	IgnoredDirectories = ".git,.idea,build,.DS_Store,.pyc,.snyk,.dccache,bin,go.mod,go.sum"
	extAsmAmd64        = "amd64"
	extAsmArm64        = "arm64"
	extCSource         = "c"
	extCHeader         = "h"
	extCppSource       = "cpp"
	extCppHeader       = "hpp"
	extGolang          = "go"
	extJson            = "json"
	extNode            = "js"
	extPython          = "py"
	extMarkdown        = "md"
	extShell           = "sh"
	extTypeScript      = "ts"
	extYaml            = "yaml"
	extYml             = "yml"

	noLinter = "has no linter"
)
