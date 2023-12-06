package monorepo

type Step struct {
	ContinueOnError bool   `yaml:"continueOnError"`
	Command         string `yaml:"command"`
}
