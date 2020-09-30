package gdobase

import (
	"database/sql"
	"github.com/jameschz/go-base/lib/gdo/cluster"
	"github.com/jameschz/go-base/lib/gdo/driver"
)

// Db :
type Db struct {
	Driver    *gdodriver.Driver   // driver
	Cluster   *gdocluster.Cluster // cluster
	Conn      *sql.DB             // connection
	Tx        *sql.Tx             // transaction
	TableName string              // table name
	TxBegin   bool                // begin tx
}

// IDb :
type IDb interface {
	Connect(driver *gdodriver.Driver) error
	T(tableName string) (db IDb)
	Close() error
	Begin() error
	Commit() error
	Rollback() error
	Max(field string) (val int64, err error)
	Min(field string) (val int64, err error)
	Shard(sql string, params ...interface{}) error
	Query(sql string, params ...interface{}) (rows *sql.Rows, err error)
	Exec(sql string, params ...interface{}) (res sql.Result, err error)
	Select(field string, where string, params ...interface{}) (rows *sql.Rows, err error)
	FetchAll(rows *sql.Rows, data interface{}) (res *[]interface{}, err error)
	FetchRow(rows *sql.Rows, data interface{}) (res interface{}, err error)
	Insert(sql string, params ...interface{}) (id int64, err error)
	Update(sql string, params ...interface{}) (affect int64, err error)
	Delete(where string, params ...interface{}) (affect int64, err error)
}
