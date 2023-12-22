Monorepo Automation
==============================

## Objectives
1. Document the `monorepo` command used for controlling the repo and automating basic functions.

Monorepo Command
---------------------------------
* The `monorepo` command is the unified tool for managing and interacting with this monorepo.
* The `monorepo` command has several sub-commands for specific functionality:
	* `clean` - remove artifacts and persistent state
	* `build` - build artifacts (binaries, etc)
	* `test` - perform unit tests
	*  `pdv` - post-deployment verification tests
	* `install` - perform local installation operations (if supported)
	* `deploy` - push artifacts to remote targets (if supported.)
	* ...more...
* The `monorepo` command can be executed from installed binary or directly using `go run go/tools/monorepo/main.go <subcommand> <options>`