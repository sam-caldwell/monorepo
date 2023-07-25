package repolinter

func (linter *Linters) Run(runner RunnerFunction, fileName *string) error {
	return runner(*fileName)
}
