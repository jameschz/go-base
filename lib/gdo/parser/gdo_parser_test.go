package gdoparser

import (
	"fmt"
	"testing"
)

func TestSqlType(t *testing.T) {
	sqlType := GetSQLType(`select * from user where 1=1`)
	fmt.Println("sqlType", sqlType)
}

func TestGetSeqIdPos(t *testing.T) {
	cs := [][]string{
		{`select * from user where id=? and name=?`, "id"},
		{`insert into user set id=?,name=?`, "name"},
		{`update user set name=? where id=?`, "id"},
	}
	for _, c := range cs {
		sqlPos := GetSeqIDPos(c[0], c[1])
		fmt.Println("sqlPos", sqlPos)
	}

}

func TestGetSeqIdVal(t *testing.T) {
	sql := `select * from user where id=? and name=?`
	val := GetSeqIDVal(sql, "id", 1, "james")
	fmt.Println("sql_val", val)
}
