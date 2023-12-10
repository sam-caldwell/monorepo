package Postgres

type ColumnType int8

const (
	Varchar ColumnType = iota
	Text
	Char
	Boolean

	reserved1
	reserved2
	reserved3
	reserved4
	reserved5
	reserved6

	Int
	Bigint
	Smallint

	reserved10
	reserved11
	reserved12
	reserved13
	reserved14
	reserved15
	reserved16

	Serial
	Serial4
	Serial8

	reserved20
	reserved21
	reserved22
	reserved23
	reserved24
	reserved25
	reserved26

	Uuid

	reserved30
	reserved31
	reserved32
	reserved33
	reserved34
	reserved35
	reserved36
)

type ColumnProperty string

const (
	PrimaryKey ColumnProperty = "primary key"
	Nullable   ColumnProperty = "null"
	NotNull    ColumnProperty = "not null"
)

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
