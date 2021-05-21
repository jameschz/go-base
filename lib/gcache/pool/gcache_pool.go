package gcachepool

import (
	"strconv"
	"strings"
	"time"

	"github.com/go-redis/redis"
	base "github.com/jameschz/go-base/lib/base"
	gcachebase "github.com/jameschz/go-base/lib/gcache/base"
	gcachedriver "github.com/jameschz/go-base/lib/gcache/driver"
	gutil "github.com/jameschz/go-base/lib/gutil"
)

var (
	_debugStatus bool
	_cPoolInit   bool
	_cPool       *base.Hmap
)

// private
func debugPrint(vals ...interface{}) {
	if _debugStatus == true {
		gutil.Dump(vals...)
	}
}

// private
func createDataSource(driver *gcachedriver.Driver, node string) *gcachebase.DataSource {
	ds := &gcachebase.DataSource{}
	ds.Name = driver.Name
	ds.Node = node
	ds.ID = gutil.UUID()
	switch driver.Type {
	case "redis":
		//add begin (db select)
		addr := node
		db := 0
		nodeArr := strings.Split(node, ":")
		if len(nodeArr) > 2 {
			var build strings.Builder
			build.WriteString(nodeArr[0])
			build.WriteString(":")
			build.WriteString(nodeArr[1])

			addr = build.String()
			db, _ = strconv.Atoi(nodeArr[2])
		}
		//add end
		ds.RedisConn = redis.NewClient(&redis.Options{
			// Basic Settings
			Addr: addr,
			DB:   db,
			// Pool Settings
			PoolSize:     driver.PoolInitSize,
			MinIdleConns: driver.PoolIdleMinSize,
			IdleTimeout:  time.Duration(driver.PoolIdleTimeoutMin * int(time.Minute)),
		})
		if ds.RedisConn == nil {
			panic("gcache> new redis client error")
		}
		_, err := ds.RedisConn.Ping().Result()
		if err != nil {
			panic("gcache> ping redis error : " + err.Error())
		}
	}
	// for debug
	debugPrint("gcachepool.createDataSource", ds)
	return ds
}

// private
func releaseDataSource(ds *gcachebase.DataSource) {
	if ds != nil {
		ds = nil
	}
	// for debug
	debugPrint("gcachepool.releaseDataSource", ds)
}

// private
func getDataSourceKey(name string, node string) string {
	return name + ":" + node
}

// SetDebug : public
func SetDebug(status bool) {
	_debugStatus = status
}

// Init : public
func Init() (err error) {
	// init once
	if _cPoolInit == true {
		return nil
	}
	// init drivers
	gcachedriver.Init()
	// init pool by drivers
	_cPool = base.NewHmap()
	cDrivers := gcachedriver.GetDrivers()
	for cName, cDriver := range cDrivers {
		for _, cNode := range cDriver.Nodes {
			cKey := getDataSourceKey(cName, cNode)
			_cPool.Set(cKey, createDataSource(cDriver, cNode))
		}
	}
	// heart beat
	go func() {
		time.Sleep(1 * time.Minute)
		for _, ds := range _cPool.Data() {
			ds.(*gcachebase.DataSource).RedisConn.Ping()
		}
	}()
	// for debug
	debugPrint("gcachepool.Init", _cPool)
	// init ok status
	if err == nil {
		_cPoolInit = true
	}
	return err
}

// Fetch : public
func Fetch(cName string, cNode string) (ds *gcachebase.DataSource, err error) {
	// get data source key
	cKey := getDataSourceKey(cName, cNode)
	// get driver by name
	ds = _cPool.Get(cKey).(*gcachebase.DataSource)
	// return ds 0
	return ds, err
}

// Return : public
func Return(ds *gcachebase.DataSource) (err error) {
	// release datasource
	releaseDataSource(ds)
	return err
}
