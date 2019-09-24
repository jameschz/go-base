package parser

import (
	"regexp"
	"strings"
)

const (
	Select = 1 << iota
	Insert
	Update
	Delete
)

func GetSqlType(sql string) int {
	sql_type := -1
	sql_arr := strings.Split(sql, " ")
	if len(sql_arr) > 0 {
		switch strings.ToUpper(sql_arr[0]) {
		case "SELECT":
			sql_type = Select
		case "INSERT":
			sql_type = Insert
		case "UPDATE":
			sql_type = Update
		case "DELETE":
			sql_type = Delete
		}
	}
	return sql_type
}

func GetSeqIdPos(sql string, seq_id string) int {
	seq_idx := -1
	var rs [][]string
	switch GetSqlType(sql) {
	//case Select:
	//case Insert:
	//case Update:
	//case Delete:
	default:
		rs = regexp.MustCompile(`([^\s,]+)=([^\s,=]+)`).FindAllStringSubmatch(sql, -1)
		//util.Dump(rs)
	}
	if len(rs) > 0 {
		for k, v := range rs {
			find := false
			for _, x := range v {
				if x == seq_id {
					find = true
				}
			}
			if find {
				seq_idx = k
				break
			}
		}
	}
	return seq_idx
}

func GetSeqIdVal(sql string, seq_id string, args ...interface{}) interface{} {
	pos := GetSeqIdPos(sql, seq_id)
	if pos >= 0 {
		return args[pos]
	}
	return nil
}
