package monorepo

type Step struct {
	Command         string                `yaml:"command"`
	Comment         string                `yaml:"comment"`
	ContinueOnError bool                  `yaml:"continueOnError"`
	Enabled         bool                  `yaml:"enabled"`
	Environment     []EnvironmentVariable `yaml:"environment"`
	ShowOutput      bool                  `yaml:"showOutput"`
	RunAsShell      bool                  `yaml:"runAsShell"`
}
