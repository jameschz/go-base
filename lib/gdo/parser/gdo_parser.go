package gdoparser

import (
	"regexp"
	"strings"
)

const (
	// Select :
	Select = 1 << iota
	// Insert :
	Insert
	// Update :
	Update
	// Delete :
	Delete
)

// GetSQLType :
func GetSQLType(sql string) int {
	sqlType := -1
	sqlArr := strings.Split(sql, " ")
	if len(sqlArr) > 0 {
		switch strings.ToUpper(sqlArr[0]) {
		case "SELECT":
			sqlType = Select
		case "INSERT":
			sqlType = Insert
		case "UPDATE":
			sqlType = Update
		case "DELETE":
			sqlType = Delete
		}
	}
	return sqlType
}

// GetSeqIDPos :
func GetSeqIDPos(sql string, seqID string) int {
	sqlIdx := -1
	var rs [][]string
	switch GetSQLType(sql) {
	//case Select:
	//case Insert:
	//case Update:
	//case Delete:
	default:
		rs = regexp.MustCompile(`([^\s,]+)=([^\s,=]+)`).FindAllStringSubmatch(sql, -1)
		//gutil.Dump(rs)
	}
	if len(rs) > 0 {
		for k, v := range rs {
			find := false
			for _, x := range v {
				if x == seqID {
					find = true
				}
			}
			if find {
				sqlIdx = k
				break
			}
		}
	}
	return sqlIdx
}

// GetSeqIDVal :
func GetSeqIDVal(sql string, seqID string, args ...interface{}) interface{} {
	pos := GetSeqIDPos(sql, seqID)
	if pos >= 0 {
		return args[pos]
	}
	return nil
}
