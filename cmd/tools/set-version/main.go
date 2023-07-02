package main

import (
	"fmt"
	"github.com/sam-caldwell/go/v2/projects/semver"
	"os"
)

const (
	versionFile = `projects/version/version.go`
	versionData = `package version

const (
	version = "%s"
)`
)

func Set(newVersion semver.SemanticVersion) (err error) {
	return os.WriteFile(versionFile, []byte(fmt.Sprintf(versionData, newVersion.String())), 0644)
}

func main() {
	var version semver.SemanticVersion
	if err := version.GetMostRecentTag(); err != nil {
		fmt.Printf("set-version failed to get current version. %s\n", err)
	}
	if err := Set(version); err != nil {
		fmt.Printf("set-version failed to set current version. %s\n", err)
	}
	fmt.Printf("Version Set: %s\n", version.String())
}
