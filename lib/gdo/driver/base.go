package driver

import (
	"go-base/lib/config"
)

type Driver struct {
	Type    string
	Host    string
	Port    string
	User    string
	Pass    string
	Charset string
	DbName  string
}

var (
	_db_drivers map[string]*Driver
)

func Init() bool {
	if len(_db_drivers) > 0 {
		return true
	}
	drivers := config.Load("database").GetStringMap("drivers")
	_db_drivers = make(map[string]*Driver, len(drivers))
	for db_name, db_conf := range drivers {
		// convert interface map
		db_driver := db_conf.(map[string]interface{})
		// check driver
		driver := &Driver{
			Type:    db_driver["type"].(string),
			Host:    db_driver["host"].(string),
			Port:    db_driver["port"].(string),
			User:    db_driver["user"].(string),
			Pass:    db_driver["pass"].(string),
			Charset: db_driver["charset"].(string),
			DbName:  db_driver["db"].(string),
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
		_db_drivers[db_name] = driver
	}
	return true
}

func GetDriver(dbs string) (driver *Driver) {
	if _, r := _db_drivers[dbs]; !r {
		panic("gdo> can not find driver")
	}
	driver = _db_drivers[dbs]
	return driver
}
