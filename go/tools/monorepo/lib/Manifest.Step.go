package monorepo

type Step struct {
	Enabled         bool   `yaml:"enabled"`
	ContinueOnError bool   `yaml:"continueOnError"`
	Command         string `yaml:"command"`
}
