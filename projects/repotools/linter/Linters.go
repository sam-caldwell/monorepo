package repolinter

type Linters struct {
	runners  map[string]LinterFunction
	commands map[string]string
}
