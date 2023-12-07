package monorepo

const (
	manifestYamlFile = "Manifest.yml"
)

type Monorepo struct {
	Root  string
	Debug bool
	//ToDo: change manifestList to an ordered list of manifests and filepaths.
	//      this will allow resorting by dependency order.
	manifestList map[string]Manifest
}
