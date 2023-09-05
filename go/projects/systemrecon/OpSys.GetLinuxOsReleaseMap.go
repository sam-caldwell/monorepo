package systemrecon

import (
	"github.com/sam-caldwell/monorepo/go/projects/keyvalue"
	words2 "github.com/sam-caldwell/monorepo/go/projects/misc/words"
	"github.com/sam-caldwell/monorepo/go/projects/wrappers/os"
)

func GetLinuxOsReleaseMap() (output keyvalue.KeyValue, err error) {
	var bytes []byte

	if bytes, err = os.ReadFile("/etc/os-release"); err != nil {
		return output, err
	}
	output.FromBytes(&bytes, words2.NewLine, words2.EqualSign)
	return output, nil
}
