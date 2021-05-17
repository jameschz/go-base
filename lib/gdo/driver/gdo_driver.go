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
	PoolMaxActive    int
	PoolMaxActiveSec int
	PoolMaxIdle      int
	PoolMaxIdleSec   int
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
			PoolMaxActive:    _dbDriver["pool_max_active"].(int),
			PoolMaxActiveSec: _dbDriver["pool_max_active_sec"].(int),
			PoolMaxIdle:      _dbDriver["pool_max_idle"].(int),
			PoolMaxIdleSec:   _dbDriver["pool_max_idle_sec"].(int),
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
func GetDriver(drName string) (driver *Driver) {
	if _, r := _dbDrivers[drName]; !r {
		panic("gdo> can not find driver")
	}
	driver = _dbDrivers[drName]
	return driver
}

// GetDrivers :
func GetDrivers() map[string]*Driver {
	return _dbDrivers
}
