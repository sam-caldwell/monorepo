package main

import (
	"fmt"
	"github.com/sam-caldwell/go/v2/projects/go/exit"
	"github.com/sam-caldwell/go/v2/projects/go/semver"
	"os"
)

const (
	versionFile = `projects/version/version.go`
	versionData = `package version

const (
	Version = "%s"
)`
)

func Set(newVersion semver.SemanticVersion) (err error) {
	return os.WriteFile(versionFile, []byte(fmt.Sprintf(versionData, newVersion.String())), 0644)
}

func main() {
	exit.IfVersionRequested()
	var version semver.SemanticVersion
	if err := version.GetMostRecentTag(); err != nil {
		fmt.Printf("set-version failed to get current version. %s\n", err)
	}
	if err := Set(version); err != nil {
		fmt.Printf("set-version failed to set current version. %s\n", err)
	}
	fmt.Printf("Version Set: %s\n", version.String())
}
