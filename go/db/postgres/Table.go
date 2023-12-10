package Postgres

type TableDescriptor struct {
	name    string
	columns Columns
}

func Table(name string, columns Columns) *TableDescriptor {
	var table TableDescriptor
	table.name = name
	table.columns = columns
	return &table
}

func (table *TableDescriptor) RenderSql() (query string) {
	//ToDo: evaluate the descriptor and render the SQL form of the table.
	return query
}
