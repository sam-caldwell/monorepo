package monorepo

// Manifest - The in-memory representation of a Manifest.yml file.
type Manifest struct {
	FileName string
	config   Config
}
