package gdomysql

import (
	"database/sql"
	"github.com/jameschz/go-base/lib/gdo/base"
	"github.com/jameschz/go-base/lib/gdo/driver"
	"github.com/jameschz/go-base/lib/gdo/parser"

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
	// open connection
	db.Driver = driver
	db.TableName = ""
	dsn := driver.User + ":" +
		driver.Pass + "@tcp(" +
		driver.Host + ":" +
		driver.Port + ")/" +
		driver.DbName + "?charset=" +
		driver.Charset
	conn, err := sql.Open("mysql", dsn)
	if err != nil {
		return err
	}
	db.Conn = conn
	return nil
}

// Close :
func (db *Mysql) Close() error {
	if db.Conn != nil {
		return db.Conn.Close()
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
		dbs := db.Cluster.Shard(params[pos].(int))
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
