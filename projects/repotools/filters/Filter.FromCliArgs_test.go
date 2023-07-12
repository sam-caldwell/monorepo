package filters

import (
	"fmt"
	"github.com/sam-caldwell/go/v2/projects/ansi"
	"os"
	"testing"

	"github.com/sam-caldwell/go/v2/projects/testing/assert"
)

func TestFromCliArgs(t *testing.T) {
	// Prepare test case
	os.Args = []string{
		"program-name",  //0
		"-commands",     //1
		"-buildenabled", //2
		"-lintenabled",  //3
		"-language",     //4
		"go",            //5
		"-os",           //6
		"linux",         //7
		"-arch",         //8
		"amd64",         //9
	}

	// Create an instance of Filter
	filter := &Filter{}

	// Call the FromCliArgs method
	err := filter.FromCliArgs()

	// Print filter settings
	ansi.Blue().Printf("Args: %v", os.Args).LF().
		Printf("Commands: %v", filter.Commands).LF().
		Printf("BuildEnabled: %v", filter.BuildEnabled).LF().
		Printf("LintEnabled: %v", filter.LintEnabled).LF().
		Printf("PackEnabled: %v", filter.PackEnabled).LF().
		Printf("ScanEnabled: %v", filter.ScanEnabled).LF().
		Printf("SignEnabled: %v", filter.SignEnabled).LF().
		Printf("Language: %v", filter.Language).LF().
		Printf("OperatingSystem: %s", filter.os).LF().
		Printf("Architecture: %s", filter.arch).LF().
		Reset()

	// Assert that no error occurred
	assert.Nil(t, err, fmt.Sprintf("Error parsing command-line arguments: '%v'", err))

	// Verify the filter settings
	assert.True(t, filter.Commands, "Commands should be enabled")
	assert.True(t, filter.BuildEnabled, "BuildEnabled should be enabled")
	assert.True(t, filter.LintEnabled, "LintEnabled should be enabled")
	assert.False(t, filter.PackEnabled, "PackEnabled should be disabled")
	assert.False(t, filter.ScanEnabled, "ScanEnabled should be disabled")
	assert.False(t, filter.SignEnabled, "SignEnabled should be disabled")
	assert.Equal(t, "linux", filter.os, "OperatingSystem should be 'linux'")
	assert.Equal(t, "amd64", filter.arch, "Architecture should be 'amd64'")
}
