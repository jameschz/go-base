package base

import (
	"go-base/lib/gdo/cluster"
	"go-base/lib/gdo/driver"
	"database/sql"
)

type Db struct {
	Driver    *driver.Driver   // driver
	Cluster   *cluster.Cluster // cluster
	Conn      *sql.DB          // connection
	Tx        *sql.Tx          // transaction
	TableName string           // table name
	TxBegin   bool             // begin tx
}

type IDb interface {
	Connect(driver *driver.Driver) error
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
	Insert(sql string, params ...interface{}) (id int64, err error)
	Update(sql string, params ...interface{}) (affect int64, err error)
	Delete(where string, params ...interface{}) (affect int64, err error)
}
