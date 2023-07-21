package repolinter

// LinterConfig - Describes a common linter function
type LinterConfig struct {
	name     string
	runner   func(filename string) error
	preCheck func() error
}
