package Postgres

type ColumnDescriptor struct {
	name         string
	typ          ColumnType
	defaultValue any
	properties   []ColumnProperty
}

func Column(name string) *ColumnDescriptor {
	var column ColumnDescriptor
	//ToDo: sanitize the name
	column.name = name
	return &column
}

func (column *ColumnDescriptor) Type(columnType ColumnType) *ColumnDescriptor {
	column.typ = columnType
	return column
}

func (column *ColumnDescriptor) Property(property ColumnProperty) *ColumnDescriptor {
	column.properties = append(column.properties, property)
	return column
}

func (column *ColumnDescriptor) Default(defaultValue string) *ColumnDescriptor {
	//ToDo: sanitize the value
	column.defaultValue = defaultValue
	return column
}
