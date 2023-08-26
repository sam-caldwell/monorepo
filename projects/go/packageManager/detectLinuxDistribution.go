package packageManager

/*
 * projects/packageManager/detectLinuxDistribution.go
 * (c) 2023 Sam Caldwell.  See LICENSE.txt
 *
 * This file will use /etc/os-release to detect the
 * Linux distribution.
 */

import (
	"github.com/sam-caldwell/go/v2/projects/go/keyvalue"
	words2 "github.com/sam-caldwell/go/v2/projects/go/misc/words"
	"github.com/sam-caldwell/go/v2/projects/go/wrappers/os"
)

// detectLinuxDistribution - detect linux distribution
func detectLinuxDistribution() (opsys string) {
	var err error
	var info keyvalue.KeyValue
	if bytes, err := os.ReadFile("/etc/os-release"); err != nil {
		return words2.EmptyString
	} else {
		info.FromBytes(&bytes, words2.NewLine, words2.EqualSign)
	}
	if opsys, err = info.GetString("ID"); err != nil {
		return words2.EmptyString
	}
	return opsys
}
