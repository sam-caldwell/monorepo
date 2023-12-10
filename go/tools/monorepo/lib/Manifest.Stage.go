package monorepo

type Stage struct {
	Comment string `yaml:"comment"`
	Enabled bool   `yaml:"enabled"`
	Steps   []Step `yaml:"steps"`
}
