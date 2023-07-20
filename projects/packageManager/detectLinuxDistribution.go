package packageManager

/*
 * projects/packageManager/detectLinuxDistribution.go
 * (c) 2023 Sam Caldwell.  See LICENSE.txt
 *
 * This file will use /etc/os-release to detect the
 * Linux distribution.
 */

import (
	keyvalue "github.com/sam-caldwell/go/v2/projects/keyvalue"
	"github.com/sam-caldwell/go/v2/projects/misc/words"
	"github.com/sam-caldwell/go/v2/projects/wrappers/os"
)

// detectLinuxDistribution - detect linux distribution
func detectLinuxDistribution() (opsys string) {
	var err error
	var info keyvalue.KeyValue
	if bytes, err := os.ReadFile("/etc/os-release"); err != nil {
		return words.EmptyString
	} else {
		info.FromBytes(&bytes, words.NewLine, words.EqualSign)
	}
	if opsys, err = info.GetString("ID"); err != nil {
		return words.EmptyString
	}
	return opsys
}
