package monorepo

const (
	manifestYamlFile = "Manifest.yml"
)

type Monorepo struct {
	Root         string
	Debug        bool
	manifestList map[string]Manifest
}
