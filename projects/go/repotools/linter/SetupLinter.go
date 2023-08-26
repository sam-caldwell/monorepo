package repolinter

func SetupLinter(useColor bool) Linters {
	var table Linters
	table.Initialize(useColor)
	return table
}
