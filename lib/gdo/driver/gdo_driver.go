package gdodriver

import (
	"github.com/jameschz/go-base/lib/config"
)

// Driver :
type Driver struct {
	Type    string
	Host    string
	Port    string
	User    string
	Pass    string
	Charset string
	DbName  string
	// pool attrs start
	PoolInitSize  int
	PoolMaxActive int
	PoolMaxIdle   int
	PoolMinIdle   int
	// pool attrs end
}

var (
	_dbDrivers map[string]*Driver
)

// Init :
func Init() bool {
	if len(_dbDrivers) > 0 {
		return true
	}
	drivers := config.Load("database").GetStringMap("drivers")
	_dbDrivers = make(map[string]*Driver, len(drivers))
	for _dbName, _dbConf := range drivers {
		// convert interface map
		_dbDriver := _dbConf.(map[string]interface{})
		// check driver
		driver := &Driver{
			Type:    _dbDriver["type"].(string),
			Host:    _dbDriver["host"].(string),
			Port:    _dbDriver["port"].(string),
			User:    _dbDriver["user"].(string),
			Pass:    _dbDriver["pass"].(string),
			Charset: _dbDriver["charset"].(string),
			DbName:  _dbDriver["db"].(string),
			// pool attrs start
			PoolInitSize:  _dbDriver["pool_init_size"].(int),
			PoolMaxActive: _dbDriver["pool_max_active"].(int),
			PoolMaxIdle:   _dbDriver["pool_max_idle"].(int),
			PoolMinIdle:   _dbDriver["pool_min_idle"].(int),
			// pool attrs end
		}
		// check driver
		if len(driver.Type) == 0 ||
			len(driver.Host) == 0 ||
			len(driver.Port) == 0 ||
			len(driver.User) == 0 ||
			len(driver.Charset) == 0 ||
			len(driver.DbName) == 0 {
			panic("gdo> init driver error")
		}
		// save driver
		_dbDrivers[_dbName] = driver
	}
	return true
}

// GetDriver :
func GetDriver(dbs string) (driver *Driver) {
	if _, r := _dbDrivers[dbs]; !r {
		panic("gdo> can not find driver")
	}
	driver = _dbDrivers[dbs]
	return driver
}

// GetDrivers :
func GetDrivers() map[string]*Driver {
	return _dbDrivers
}
