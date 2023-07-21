package repolinter

func SetupLinter() Linters {
	var table Linters
	table.Initialize()
	return table
}
