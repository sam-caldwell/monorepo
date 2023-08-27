package repolinter

// LinterConfig - Describes a common linter function
type LinterConfig struct {
	enabled        bool
	name     string
	runner   RunnerFunction
	preCheck func() error
	directoryLevel bool
}

type RunnerFunction func(filename string) error
