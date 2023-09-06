package systemrecon

import (
	"github.com/sam-caldwell/monorepo/go/keyvalue"
	"github.com/sam-caldwell/monorepo/go/misc/words"
	"github.com/sam-caldwell/monorepo/go/wrappers/os"
)

func GetLinuxOsReleaseMap() (output keyvalue.KeyValue, err error) {
	var bytes []byte

	if bytes, err = os.ReadFile("/etc/os-release"); err != nil {
		return output, err
	}
	output.FromBytes(&bytes, words.NewLine, words.EqualSign)
	return output, nil
}
