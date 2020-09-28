package gdo

import (
	"github.com/jameschz/go-base/lib/gdo/base"
	"github.com/jameschz/go-base/lib/gdo/cluster"
	"github.com/jameschz/go-base/lib/gdo/driver"
	"github.com/jameschz/go-base/lib/gdo/mysql"
)

// D : connect by driver
func D(dbs string) (idb gdobase.IDb) {
	// init driver
	gdodriver.Init()
	// get db driver
	dbDriver := gdodriver.GetDriver(dbs)
	if len(dbDriver.Type) == 0 {
		panic("gdo> db driver error")
	}
	// db initialize
	var err error
	switch dbDriver.Type {
	case "mysql":
		mysql := &gdomysql.Mysql{}
		err = mysql.Connect(dbDriver)
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

// C : connect by cluster
func C(cs string) (idb gdobase.IDb) {
	// init all
	gdocluster.Init()
	// get db cluster
	dbCluster := gdocluster.GetCluster(cs)
	if len(dbCluster.Type) == 0 {
		panic("gdo> db cluster error")
	}
	// db initialize
	switch dbCluster.Type {
	case "mysql":
		mysql := &gdomysql.Mysql{}
		mysql.Cluster = dbCluster
		idb = mysql
	default:
		panic("gdo> unknown driver type")
	}
	return idb
}
