package systemrecon

import (
	keyvalue "github.com/sam-caldwell/go/v2/projects/KeyValue"
	"github.com/sam-caldwell/go/v2/projects/misc/words"
	"github.com/sam-caldwell/go/v2/projects/wrappers/os"
)

func GetLinuxOsReleaseMap() (output keyvalue.KeyValue, err error) {
	var bytes []byte

	if bytes, err = os.ReadFile("/etc/os-release"); err != nil {
		return output, err
	}
	output.FromBytes(&bytes, words.NewLine, words.EqualSign)
	return output, nil
}
