package systemrecon

import (
	"github.com/sam-caldwell/monorepo/projects/keyvalue"
	"github.com/sam-caldwell/monorepo/projects/misc/words"
	"github.com/sam-caldwell/monorepo/projects/wrappers/os"
)

func GetLinuxOsReleaseMap() (output keyvalue.KeyValue[string, string], err error) {
	var bytes []byte

	if bytes, err = os.ReadFile("/etc/os-release"); err != nil {
		return output, err
	}
	output.FromBytes(&bytes, words.NewLine, words.EqualSign)
	return output, nil
}
