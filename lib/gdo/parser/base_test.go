package parser

import (
	"fmt"
	"testing"
)

func TestSqlType(t *testing.T) {
	sql_type := GetSqlType(`select * from user where 1=1`)
	fmt.Println("sql_type", sql_type)
}

func TestGetSeqIdPos(t *testing.T) {
	cs := [][]string{
		[]string{`select * from user where id=? and name=?`, "id"},
		[]string{`insert into user set id=?,name=?`, "name"},
		[]string{`update user set name=? where id=?`, "id"},
	}
	for _, c := range cs {
		sql_pos := GetSeqIdPos(c[0], c[1])
		fmt.Println("sql_pos", sql_pos)
	}

}

func TestGetSeqIdVal(t *testing.T) {
	sql := `select * from user where id=? and name=?`
	val := GetSeqIdVal(sql, "id", 1, "james")
	fmt.Println("sql_val", val)
}
