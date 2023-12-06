package monorepo

type Stage struct {
	Enabled bool   `yaml:"enabled"`
	Steps   []Step `yaml:"steps"`
}
