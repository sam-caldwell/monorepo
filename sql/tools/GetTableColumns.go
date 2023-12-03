package database

import (
	"fmt"
	Postgres "github.com/sam-caldwell/monorepo/go/db/postgres"
	"testing"
)

// GetTableColumns - Enumerate the table columns and properties for test evaluation.
func GetTableColumns(t *testing.T, db *Postgres.Db, tableName string) (columns []string) {
	t.Logf("tableName: %s", tableName)
	query := fmt.Sprintf(""+
		"SELECT "+
		"coalesce(column_name,'<<null>>') as ColumnName, "+
		"coalesce(data_type,'<<null>>') as DataType, "+
		"coalesce(is_nullable,'<<null>>') as IsNullable, "+
		"coalesce(column_default,'<<null>>') as ColumnDefault "+
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
		columns = append(columns,
			fmt.Sprintf(""+
				"ColumnName:%v,"+
				"DataType:%v,"+
				"IsNullable:%v,"+
				"ColumnDefault:%v",
				columnName, dataType, isNullable, columnDefault))
	}
	return columns
}
