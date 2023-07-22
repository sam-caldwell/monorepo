package main

import (
	"github.com/sam-caldwell/go/v2/projects/ansi"
	"github.com/sam-caldwell/go/v2/projects/exit"
	"os"
)

const preComment = `
# git/hooks/pre-commit
# (c) 2023 Sam Caldwell.  See LICENSE.txt
# This script will run the lint-repo golang command
go run cmd/tools/lint-repo/main.go -color -quiet || exit 1
`

const prePush = `
# git/hooks/pre-push
# (c) 2023 Sam Caldwell.  See LICENSE.txt
# This script will run the tests, security scanners and builds
go run cmd/tools/run-tests/main.go -color -quiet || exit 1
go run cmd/tools/run-builds/main.go -color -quiet || exit 1
go run cmd/tools/run-scans/main.go -color -quiet || exit 1
`

const commandUsage = `
update-git-hooks [-h|--help]
	show this screen

update-git-hooks
	install the git hooks scripts
`

func main() {
	const program = "update-git-hooks"

	exit.IfHelpRequested(commandUsage)
	exit.IfVersionRequested()

	hookScripts := map[string]string{
		".git/hooks/pre-commit": preComment,
		".git/hooks/pre-push":   prePush,
	}

	for file, script := range hookScripts {
		if err := os.WriteFile(file, []byte(script), 0755); err != nil {
			ansi.Red().Printf("[%s]FAIL on %s", program, file).LF().Reset()
		}
		ansi.Green().Printf("[%s]PASS on %s", program, file).LF().Reset()
	}

}