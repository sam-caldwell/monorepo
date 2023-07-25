package repolinter

import (
	"fmt"
	"github.com/sam-caldwell/go/v2/projects/exit/errors"
	"github.com/sam-caldwell/go/v2/projects/fs/file"
	keyvalue "github.com/sam-caldwell/go/v2/projects/keyvalue"
)

func LoadProjectManifest(manifest string) (kv keyvalue.KeyValue, err error) {
	if !file.Exists(manifest) {
		return kv, fmt.Errorf(errors.NotFound+errors.Details, manifest)
	}
	if err := kv.FromFile(manifest, ":", "\n"); err != nil {
		return kv, err
	}
	return kv, err
}
