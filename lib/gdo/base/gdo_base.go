package gdobase

import (
	"database/sql"

	gdocluster "github.com/jameschz/go-base/lib/gdo/cluster"
	gdodriver "github.com/jameschz/go-base/lib/gdo/driver"
)

// DataSource :
type DataSource struct {
	ID   string
	Name string
	Conn *sql.DB
}

// Db :
type Db struct {
	Driver     *gdodriver.Driver   // driver
	Cluster    *gdocluster.Cluster // cluster
	DataSource *DataSource         // datasource
	Conn       *sql.DB             // db connection
	Tx         *sql.Tx             // transaction
	TxBegin    bool                // begin tx
	TableName  string              // table name
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
	FetchStructs(rows *sql.Rows, data interface{}) (res *[]interface{}, err error)
	FetchStruct(rows *sql.Rows, data interface{}) (res interface{}, err error)
	FetchMaps(rows *sql.Rows) (res *[]map[string]interface{}, err error)
	FetchMap(rows *sql.Rows) (res map[string]interface{}, err error)
	Insert(sql string, params ...interface{}) (id int64, err error)
	Update(sql string, params ...interface{}) (affect int64, err error)
	Delete(where string, params ...interface{}) (affect int64, err error)
}
