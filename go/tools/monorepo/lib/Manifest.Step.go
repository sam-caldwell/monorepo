package monorepo

type Step struct {
	ContinueOnError bool   `yaml:"continueOnError"`
	Enabled         bool   `yaml:"enabled"`
	Command         string `yaml:"command"`
}
