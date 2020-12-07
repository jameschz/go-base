package gcachepool

import (
	"errors"
	"strconv"
	"strings"
	"sync"

	"github.com/go-redis/redis"
	base "github.com/jameschz/go-base/lib/base"
	gcachebase "github.com/jameschz/go-base/lib/gcache/base"
	gcachedriver "github.com/jameschz/go-base/lib/gcache/driver"
	gutil "github.com/jameschz/go-base/lib/gutil"
)

var (
	_debugStatus bool
	_cPoolInit   bool
	_cPoolIdle   map[string]*base.Stack
	_cPoolActive map[string]*base.Hmap
	_cPoolLock   sync.Mutex
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
			// Addr: node,
			Addr: addr,
			DB:   db,
		})
	}
	// for debug
	debugPrint("gcachepool.createDataSource", ds)
	return ds
}

// private
func releaseDataSource(ds *gcachebase.DataSource) {
	if ds.RedisConn != nil {
		ds.RedisConn.Close()
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
	cDrivers := gcachedriver.GetDrivers()
	_cPoolIdle = make(map[string]*base.Stack, 0)
	_cPoolActive = make(map[string]*base.Hmap, 0)
	for cName, cDriver := range cDrivers {
		for _, cNode := range cDriver.Nodes {
			cKey := getDataSourceKey(cName, cNode)
			_cPoolIdle[cKey] = base.NewStack()
			_cPoolActive[cKey] = base.NewHmap()
			for i := 0; i < cDriver.PoolInitSize; i++ {
				_cPoolIdle[cKey].Push(createDataSource(cDriver, cNode))
			}
		}
	}
	// for debug
	debugPrint("gcachepool.Init", _cPoolIdle, _cPoolActive)
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
	cDriver := gcachedriver.GetDriver(cName)
	// fetch start >>> lock
	_cPoolLock.Lock()
	// reach to max active size
	activeSize := _cPoolActive[cKey].Len()
	if cDriver.PoolMaxActive <= activeSize {
		return nil, errors.New("gcachepool : max active limit")
	}
	// add if not enough
	idleSize := _cPoolIdle[cKey].Len()
	if cDriver.PoolMinIdle >= idleSize {
		idleSizeAdd := cDriver.PoolMaxIdle - idleSize
		for i := 0; i < idleSizeAdd; i++ {
			_cPoolIdle[cKey].Push(createDataSource(cDriver, cNode))
		}
		// for debug
		debugPrint("gcachepool.Fetch Add", _cPoolIdle[cKey].Len(), _cPoolActive[cKey].Len())
	}
	// fetch from front
	if _cPoolIdle[cKey].Len() >= 1 {
		ds = _cPoolIdle[cKey].Pop().(*gcachebase.DataSource)
		_cPoolActive[cKey].Set(ds.ID, ds)
	} else {
		return nil, errors.New("gcachepool : no enough ds")
	}
	// for debug
	debugPrint("gcachepool.Fetch", _cPoolIdle[cKey].Len(), _cPoolActive[cKey].Len())
	// fetch end >>> unlock
	_cPoolLock.Unlock()
	// return ds 0
	return ds, err
}

// Return : public
func Return(ds *gcachebase.DataSource) (err error) {
	// get data source key
	cKey := getDataSourceKey(ds.Name, ds.Node)
	// get driver by name
	cDriver := gcachedriver.GetDriver(ds.Name)
	// return start >>> lock
	_cPoolLock.Lock()
	// delete from active list
	_cPoolActive[cKey].Delete(ds.ID)
	// return or release
	idleSize := _cPoolIdle[cKey].Len()
	if cDriver.PoolMaxIdle <= idleSize {
		releaseDataSource(ds)
	} else {
		_cPoolIdle[cKey].Push(ds)
	}
	// return end >>> unlock
	_cPoolLock.Unlock()
	// for debug
	debugPrint("gcachepool.Return", _cPoolIdle[cKey].Len(), _cPoolActive[cKey].Len())
	return err
}
