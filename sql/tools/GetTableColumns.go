package database

import (
	"fmt"
	Postgres "github.com/sam-caldwell/monorepo/go/db/postgres"
	"testing"
)

type Record struct {
	ColumnName    string
	DataType      string
	IsNullable    string
	ColumnDefault string
}

// GetTableColumns - Enumerate the table columns and properties for test evaluation.
func GetTableColumns(t *testing.T, db *Postgres.Db, tableName string) (columns []Record) {

	query := fmt.Sprintf(""+
		"SELECT "+
		"coalesce(column_name,'<<null>>'), "+
		"coalesce(data_type,'<<null>>'), "+
		"coalesce(is_nullable,'<<null>>'), "+
		"coalesce(column_default,'<<null>>') "+
		"FROM information_schema.columns "+
		"WHERE table_name = '%s';",
		tableName)

	rows, err := db.Query(query)
	CheckError(t, err)
	defer func() {
		CheckError(t, rows.Close())
	}()

	for rows.Next() {
		var columnName, dataType, isNullable, columnDefault string
		CheckError(t, rows.Scan(&columnName, &dataType, &isNullable, &columnDefault))
		columns = append(columns, Record{
			ColumnName:    fmt.Sprintf("%v", columnName),
			DataType:      fmt.Sprintf("%v", dataType),
			IsNullable:    fmt.Sprintf("%v", isNullable),
			ColumnDefault: fmt.Sprintf("%v", columnDefault),
		})
	}
	return columns

}
