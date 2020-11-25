package gdomysql

import (
	"database/sql"
	"reflect"

	gdobase "github.com/jameschz/go-base/lib/gdo/base"
	gdodriver "github.com/jameschz/go-base/lib/gdo/driver"
	gdoparser "github.com/jameschz/go-base/lib/gdo/parser"
	gdopool "github.com/jameschz/go-base/lib/gdo/pool"

	// import mysql lib
	_ "github.com/go-sql-driver/mysql"
)

// Mysql :
type Mysql struct {
	gdobase.Db
}

// T :
func (db *Mysql) T(table string) gdobase.IDb {
	db.TableName = table
	return db
}

// GetTable :
func (db *Mysql) GetTable() (string, error) {
	if len(db.TableName) == 0 {
		panic("gdo.mysql> can not find table")
	}
	return db.TableName, nil
}

// CleanTable :
func (db *Mysql) CleanTable() *Mysql {
	db.TableName = ""
	return db
}

// Connect :
func (db *Mysql) Connect(driver *gdodriver.Driver) error {
	// connect once
	if db.Conn == nil {
		// gdopool
		gdopool.Init()
		dataSource, err := gdopool.Fetch(driver.DbName)
		// init db vars
		db.Driver = driver
		db.DataSource = dataSource
		db.Conn = dataSource.Conn
		db.TableName = ""
		return err
	}
	return nil
}

// Close :
func (db *Mysql) Close() error {
	// close once
	if db.Conn != nil {
		err := gdopool.Return(db.DataSource)
		db.DataSource = nil
		db.Conn = nil
		db.Tx = nil
		return err
	}
	return nil
}

// Begin :
func (db *Mysql) Begin() error {
	// see begin tx logic in Shard()
	db.TxBegin = true
	return nil
}

// Commit :
func (db *Mysql) Commit() error {
	// commit tx logic
	if db.Tx != nil {
		if err := db.Tx.Commit(); err != nil {
			return err
		}
		db.Tx = nil
		db.TxBegin = false
	}
	return nil
}

// Rollback :
func (db *Mysql) Rollback() error {
	// rollback tx logic
	if db.Tx != nil {
		if err := db.Tx.Rollback(); err != nil {
			return err
		}
		db.Tx = nil
		db.TxBegin = false
	}
	return nil
}

// Shard :
func (db *Mysql) Shard(sqlStr string, params ...interface{}) error {
	// connect for sharding
	if db.Conn == nil {
		pos := gdoparser.GetSeqIDPos(sqlStr, db.Cluster.SeqID)
		if pos < 0 {
			panic("gdo> can not find seq value")
		}
		dbs := db.Cluster.Shard(params[pos].(int64))
		driver := gdodriver.GetDriver(dbs)
		if len(driver.Type) == 0 {
			panic("gdo> db driver error")
		}
		if err := db.Connect(driver); err != nil {
			return err
		}
	}
	// begin tx logic
	if db.TxBegin == true {
		tx, err := db.Conn.Begin()
		if err != nil {
			return err
		}
		db.Tx = tx
	}
	return nil
}

// Query :
func (db *Mysql) Query(sqlStr string, params ...interface{}) (rows *sql.Rows, err error) {
	// do sharding
	if err = db.Shard(sqlStr, params...); err != nil {
		return nil, err
	}
	// do prepare, don't prepare in tx for buffer error
	var stmt *sql.Stmt
	stmt, err = db.Conn.Prepare(sqlStr)
	if err != nil {
		return nil, err
	}
	// do query
	defer stmt.Close()
	if rows, err = stmt.Query(params...); err != nil {
		return nil, err
	}
	return rows, nil
}

// Exec :
func (db *Mysql) Exec(sqlStr string, params ...interface{}) (res sql.Result, err error) {
	// do sharding
	if err = db.Shard(sqlStr, params...); err != nil {
		return nil, err
	}
	// do prepare
	var stmt *sql.Stmt
	if db.Tx != nil {
		stmt, err = db.Tx.Prepare(sqlStr)
	} else {
		stmt, err = db.Conn.Prepare(sqlStr)
	}
	if err != nil {
		return nil, err
	}
	// do exec
	defer stmt.Close()
	if res, err = stmt.Exec(params...); err != nil {
		return res, err
	}
	return res, nil
}

// Max :
func (db *Mysql) Max(field string) (val int64, err error) {
	table, err := db.GetTable()
	defer db.CleanTable()
	if err != nil {
		return -1, err
	}
	sql := "select max(" + field + ") from " + table
	rows, err := db.Query(sql)
	if err != nil {
		return -1, err
	}
	for rows.Next() {
		rows.Scan(&val)
	}
	return val, nil
}

// Min :
func (db *Mysql) Min(field string) (val int64, err error) {
	table, err := db.GetTable()
	defer db.CleanTable()
	if err != nil {
		return -1, err
	}
	sql := "select min(" + field + ") from " + table
	rows, err := db.Query(sql)
	if err != nil {
		return -1, err
	}
	for rows.Next() {
		rows.Scan(&val)
	}
	return val, nil
}

// Select :
func (db *Mysql) Select(field string, where string, params ...interface{}) (rows *sql.Rows, err error) {
	table, err := db.GetTable()
	defer db.CleanTable()
	if err != nil {
		return nil, err
	}
	sql := "select " + field + " from " + table + " where " + where
	rows, err = db.Query(sql, params...)
	if err != nil {
		return nil, err
	}
	return rows, nil
}

// FetchStructs :
func (db *Mysql) FetchStructs(rows *sql.Rows, data interface{}) (*[]interface{}, error) {
	var err error
	stu := reflect.ValueOf(data).Elem()
	len := stu.NumField()
	list := make([]interface{}, 0)
	cols := make([]interface{}, len)
	for i := 0; i < len; i++ {
		cols[i] = stu.Field(i).Addr().Interface()
	}
	// scan all rows into cols
	for rows.Next() {
		err = rows.Scan(cols...)
		if err != nil {
			break
		}
		// add item into list
		list = append(list, stu.Interface())
	}
	// close rows
	_ = rows.Close()
	// return list
	return &list, err
}

// FetchStruct :
func (db *Mysql) FetchStruct(rows *sql.Rows, data interface{}) (interface{}, error) {
	res, err := db.FetchStructs(rows, data)
	for _, v := range *res {
		return v, err
	}
	return nil, nil
}

// FetchMaps :
func (db *Mysql) FetchMaps(rows *sql.Rows) (*[]map[string]interface{}, error) {
	var err error
	// init cols cache
	columns, _ := rows.Columns()
	columnLen := len(columns)
	cols := make([]interface{}, columnLen)
	for i := range cols {
		var a interface{}
		cols[i] = &a
	}
	var list []map[string]interface{}
	// scan all rows into cols
	for rows.Next() {
		err = rows.Scan(cols...)
		if err != nil {
			break
		}
		// build item data
		item := make(map[string]interface{})
		for i, v := range cols {
			c := *v.(*interface{})
			switch c.(type) {
			case []uint8:
				item[columns[i]] = string(c.([]uint8))
				break
			case nil:
				item[columns[i]] = ""
				break
			default:
				item[columns[i]] = c
				break
			}
		}
		// add item into list
		list = append(list, item)
	}
	// close rows
	_ = rows.Close()
	// return list
	return &list, err
}

// FetchMap :
func (db *Mysql) FetchMap(rows *sql.Rows) (map[string]interface{}, error) {
	res, err := db.FetchMaps(rows)
	for _, v := range *res {
		return v, err
	}
	return nil, nil
}

// Insert :
func (db *Mysql) Insert(insert string, params ...interface{}) (id int64, err error) {
	table, err := db.GetTable()
	defer db.CleanTable()
	if err != nil {
		return -1, err
	}
	sql := "insert " + table + " set " + insert
	res, err := db.Exec(sql, params...)
	if err != nil {
		return 0, err
	}
	id, err = res.LastInsertId()
	return id, err
}

// Update :
func (db *Mysql) Update(update string, params ...interface{}) (affect int64, err error) {
	table, err := db.GetTable()
	defer db.CleanTable()
	if err != nil {
		return -1, err
	}
	sql := "update " + table + " set " + update
	res, err := db.Exec(sql, params...)
	if err != nil {
		return 0, err
	}
	affect, err = res.RowsAffected()
	return affect, err
}

// Delete :
func (db *Mysql) Delete(where string, params ...interface{}) (affect int64, err error) {
	table, err := db.GetTable()
	defer db.CleanTable()
	if err != nil {
		return -1, err
	}
	sql := "delete from " + table + " where " + where
	res, err := db.Exec(sql, params...)
	if err != nil {
		return 0, err
	}
	affect, err = res.RowsAffected()
	return affect, err
}
