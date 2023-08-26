package repolinter

import (
	"fmt"
	"github.com/sam-caldwell/go/v2/projects/go/exit/errors"
	"github.com/sam-caldwell/go/v2/projects/go/fs/file"
	"github.com/sam-caldwell/go/v2/projects/go/keyvalue"
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
