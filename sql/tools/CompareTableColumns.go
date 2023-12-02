package database

import (
	"testing"
)

func CompareTableColumns(t *testing.T, lhs []Record, rhs []Record) {
	for i := 0; i < len(lhs); i++ {
		if lhs[i].ColumnName != rhs[i].ColumnName {
			t.Fatalf("ColumnName mismatch(%d): %s", i, lhs[i].ColumnName)
		}
		if lhs[i].ColumnDefault != rhs[i].ColumnDefault {
			t.Fatalf("ColumnDefault mismatch(%d): %s", i, lhs[i].ColumnDefault)
		}
		if lhs[i].DataType != rhs[i].DataType {
			t.Fatalf("DataType mismatch(%d): %s", i, lhs[i].DataType)
		}
		if lhs[i].IsNullable != rhs[i].IsNullable {
			t.Fatalf("IsNullable mismatch(%d): %s", i, lhs[i].IsNullable)
		}
	}
}
