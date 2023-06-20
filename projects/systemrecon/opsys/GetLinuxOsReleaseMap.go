package systemrecon

import (
	misc "github.com/sam-caldwell/go/v2/projects/misc/formatting"
	"github.com/sam-caldwell/go/v2/projects/misc/words"
	"os"
)

func GetLinuxOsReleaseMap() (output map[string]string, err error) {
	var bytes []byte

	if bytes, err = os.ReadFile("/etc/os-release"); err != nil {
		return nil, err
	}
	return misc.TextToKeyValueMap(&bytes, words.NewLine, words.EqualSign, false), nil
}
