package gdo

import (
	"go-base/lib/gdo/base"
	"go-base/lib/gdo/cluster"
	"go-base/lib/gdo/driver"
	"go-base/lib/gdo/mysql"
)

// connect by driver
func D(dbs string) (idb base.IDb) {
	// init driver
	driver.Init()
	// get db driver
	db_driver := driver.GetDriver(dbs)
	if len(db_driver.Type) == 0 {
		panic("gdo> db driver error")
	}
	// db initialize
	var err error
	switch db_driver.Type {
	case "mysql":
		mysql := &mysql.Mysql{}
		err = mysql.Connect(db_driver)
		idb = mysql
	default:
		panic("gdo> unknown driver type")
	}
	// throw error
	if err != nil {
		panic("gdo> db connect error")
	}
	return idb
}

// connect by cluster
func C(cs string) (idb base.IDb) {
	// init all
	cluster.Init()
	// get db cluster
	db_cluster := cluster.GetCluster(cs)
	if len(db_cluster.Type) == 0 {
		panic("gdo> db cluster error")
	}
	// db initialize
	switch db_cluster.Type {
	case "mysql":
		mysql := &mysql.Mysql{}
		mysql.Cluster = db_cluster
		idb = mysql
	default:
		panic("gdo> unknown driver type")
	}
	return idb
}
