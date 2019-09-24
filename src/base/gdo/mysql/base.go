package mysql

import (
	"base/gdo/base"
	"base/gdo/driver"
	"base/gdo/parser"
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

type Mysql struct {
	base.Db
}

func (db *Mysql) T(table string) base.IDb {
	db.TableName = table
	return db
}

func (db *Mysql) GetTable() (string, error) {
	if len(db.TableName) == 0 {
		panic("gdo.mysql> can not find table")
	}
	return db.TableName, nil
}

func (db *Mysql) CleanTable() *Mysql {
	db.TableName = ""
	return db
}

func (db *Mysql) Connect(driver *driver.Driver) error {
	// open connection
	db.Driver = driver
	db.TableName = ""
	dsn := driver.User + ":" +
		driver.Pass + "@tcp(" +
		driver.Host + ":" +
		driver.Port + ")/" +
		driver.DbName + "?charset=" +
		driver.Charset
	if conn, err := sql.Open("mysql", dsn); err != nil {
		return err
	} else {
		db.Conn = conn
	}
	return nil
}

func (db *Mysql) Close() error {
	if db.Conn != nil {
		return db.Conn.Close()
	}
	return nil
}

func (db *Mysql) Begin() error {
	// see begin tx logic in Shard()
	db.TxBegin = true
	return nil
}

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

func (db *Mysql) Shard(sql_s string, params ...interface{}) error {
	// connect for sharding
	if db.Conn == nil {
		pos := parser.GetSeqIdPos(sql_s, db.Cluster.SeqId)
		if pos < 0 {
			panic("gdo> can not find seq value")
		}
		dbs := db.Cluster.Shard(params[pos].(int))
		driver := driver.GetDriver(dbs)
		if len(driver.Type) == 0 {
			panic("gdo> db driver error")
		}
		if err := db.Connect(driver); err != nil {
			return err
		}
	}
	// begin tx logic
	if db.TxBegin == true {
		if tx, err := db.Conn.Begin(); err != nil {
			return err
		} else {
			db.Tx = tx
		}
	}
	return nil
}

func (db *Mysql) Query(sql_s string, params ...interface{}) (rows *sql.Rows, err error) {
	// do sharding
	if err = db.Shard(sql_s, params...); err != nil {
		return nil, err
	}
	// do prepare, don't prepare in tx for buffer error
	var stmt *sql.Stmt
	stmt, err = db.Conn.Prepare(sql_s)
	if err != nil {
		return nil, err
	}
	// do query
	defer stmt.Close()
	if rows, err = stmt.Query(params...); err != nil {
		return nil, err
	} else {
		return rows, nil
	}
}

func (db *Mysql) Exec(sql_s string, params ...interface{}) (res sql.Result, err error) {
	// do sharding
	if err = db.Shard(sql_s, params...); err != nil {
		return nil, err
	}
	// do prepare
	var stmt *sql.Stmt
	if db.Tx != nil {
		stmt, err = db.Tx.Prepare(sql_s)
	} else {
		stmt, err = db.Conn.Prepare(sql_s)
	}
	if err != nil {
		return nil, err
	}
	// do exec
	defer stmt.Close()
	if res, err = stmt.Exec(params...); err != nil {
		return res, err
	} else {
		return res, nil
	}
}

func (db *Mysql) Max(field string) (val int64, err error) {
	table, err := db.GetTable()
	defer db.CleanTable()
	if err != nil {
		return -1, err
	}
	sql := "select max(" + field + ") from " + table
	if rows, err := db.Query(sql); err != nil {
		return -1, err
	} else {
		for rows.Next() {
			rows.Scan(&val)
		}
		return val, nil
	}
}

func (db *Mysql) Min(field string) (val int64, err error) {
	table, err := db.GetTable()
	defer db.CleanTable()
	if err != nil {
		return -1, err
	}
	sql := "select max(" + field + ") from " + table
	if rows, err := db.Query(sql); err != nil {
		return -1, err
	} else {
		for rows.Next() {
			rows.Scan(&val)
		}
		return val, nil
	}
}

func (db *Mysql) Select(field string, where string, params ...interface{}) (rows *sql.Rows, err error) {
	table, err := db.GetTable()
	defer db.CleanTable()
	if err != nil {
		return nil, err
	}
	sql := "select " + field + " from " + table + " where " + where
	if rows, err := db.Query(sql, params...); err != nil {
		return nil, err
	} else {
		return rows, nil
	}
}

func (db *Mysql) Insert(insert string, params ...interface{}) (id int64, err error) {
	table, err := db.GetTable()
	defer db.CleanTable()
	if err != nil {
		return -1, err
	}
	sql := "insert " + table + " set " + insert
	if res, err := db.Exec(sql, params...); err != nil {
		return 0, err
	} else {
		id, err = res.LastInsertId()
		return id, err
	}
}

func (db *Mysql) Update(update string, params ...interface{}) (affect int64, err error) {
	table, err := db.GetTable()
	defer db.CleanTable()
	if err != nil {
		return -1, err
	}
	sql := "update " + table + " set " + update
	if res, err := db.Exec(sql, params...); err != nil {
		return 0, err
	} else {
		affect, err = res.RowsAffected()
		return affect, err
	}
}

func (db *Mysql) Delete(where string, params ...interface{}) (affect int64, err error) {
	table, err := db.GetTable()
	defer db.CleanTable()
	if err != nil {
		return -1, err
	}
	sql := "delete from " + table + " where " + where
	if res, err := db.Exec(sql, params...); err != nil {
		return 0, err
	} else {
		affect, err = res.RowsAffected()
		return affect, err
	}
}
