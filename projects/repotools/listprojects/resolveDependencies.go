package listprojects

import (
	"fmt"
	keyvalue "github.com/sam-caldwell/go/v2/projects/keyvalue"
	"github.com/sam-caldwell/go/v2/projects/keyvalue/pair"
	projectmanifest "github.com/sam-caldwell/go/v2/projects/repotools/manifest"
)

func resolveDependencies(projects *keyvalue.KeyValue, list *pair.OrderedPair,
	thisProject *string, manifest *projectmanifest.Manifest) (err error) {

	for _, dependency := range manifest.Dependencies {
		if !list.Has(dependency) {
			var depManifest projectmanifest.Manifest
			if raw, found := projects.FindKey(dependency); !found {
				return fmt.Errorf("dependency not found: %s", dependency)
			} else {
				depManifest = raw.(projectmanifest.Manifest)
			}
			err = resolveDependencies(projects, list, &dependency, &depManifest)
			if err != nil {
				return err
			}
		}
	}
	list.Add(*thisProject, *manifest)
	return nil
}
