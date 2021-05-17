package gdopool

import (
	"database/sql"
	"time"

	base "github.com/jameschz/go-base/lib/base"
	gdobase "github.com/jameschz/go-base/lib/gdo/base"
	gdodriver "github.com/jameschz/go-base/lib/gdo/driver"
	gutil "github.com/jameschz/go-base/lib/gutil"

	// import mysql lib
	_ "github.com/go-sql-driver/mysql"
)

var (
	_debugStatus bool
	_dbPoolInit  bool
	_dbPool      *base.Hmap
)

// private
func debugPrint(vals ...interface{}) {
	if _debugStatus {
		gutil.Dump(vals...)
	}
}

// private
func createDataSource(driver *gdodriver.Driver) *gdobase.DataSource {
	ds := &gdobase.DataSource{}
	ds.Name = driver.DbName
	ds.ID = gutil.UUID()
	// open db connection
	dsn := driver.User + ":" +
		driver.Pass + "@tcp(" +
		driver.Host + ":" +
		driver.Port + ")/" +
		driver.DbName + "?charset=" +
		driver.Charset
	switch driver.Type {
	case "mysql":
		dbc, err := sql.Open("mysql", dsn)
		if err != nil {
			panic("gdo> open db error")
		}
		err = dbc.Ping()
		if err != nil {
			panic("gdo> ping db error : " + err.Error())
		}
		dbc.SetMaxOpenConns(driver.PoolMaxActive)
		dbc.SetMaxIdleConns(driver.PoolMaxIdle)
		dbc.SetConnMaxLifetime(time.Duration(driver.PoolMaxActiveSec * int(time.Second)))
		dbc.SetConnMaxIdleTime(time.Duration(driver.PoolMaxIdleSec * int(time.Second)))
		ds.Conn = dbc
	}
	// for debug
	debugPrint("gdopool.createDataSource", ds)
	return ds
}

// private
func releaseDataSource(ds *gdobase.DataSource) {
	if ds != nil {
		ds = nil
	}
	// for debug
	debugPrint("gdopool.releaseDataSource", ds)
}

// SetDebug : public
func SetDebug(status bool) {
	_debugStatus = status
}

// Init : public
func Init() (err error) {
	// init once
	if _dbPoolInit {
		return nil
	}
	// init drivers
	gdodriver.Init()
	// init pool by drivers
	_dbPool = base.NewHmap()
	dbDrivers := gdodriver.GetDrivers()
	for _, dbDriver := range dbDrivers {
		_dbPool.Set(dbDriver.DbName, createDataSource(dbDriver))
	}
	// for debug
	debugPrint("gdopool.Init", _dbPool)
	// init ok status
	if err == nil {
		_dbPoolInit = true
	}
	return err
}

// Fetch : public
func Fetch(dbName string) (ds *gdobase.DataSource, err error) {
	// get datasource from pool
	ds = _dbPool.Get(dbName).(*gdobase.DataSource)
	// return ds 0
	return ds, err
}

// Return : public
func Return(ds *gdobase.DataSource) (err error) {
	// return start >>> lock
	releaseDataSource(ds)
	// return 0
	return err
}
