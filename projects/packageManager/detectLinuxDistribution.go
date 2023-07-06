package packageManager

import (
	keyvalue "github.com/sam-caldwell/go/v2/projects/KeyValue"
	"github.com/sam-caldwell/go/v2/projects/misc/words"
	"os"
)

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
