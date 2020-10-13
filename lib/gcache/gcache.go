package gcache

import (
	gcachebase "github.com/jameschz/go-base/lib/gcache/base"
	gcachedriver "github.com/jameschz/go-base/lib/gcache/driver"
	gcacheredis "github.com/jameschz/go-base/lib/gcache/redis"
	gcacheregion "github.com/jameschz/go-base/lib/gcache/region"
)

// D : connect by driver
func D(cs string) (ic gcachebase.ICache) {
	// init driver
	gcachedriver.Init()
	// get mq driver
	_cDriver := gcachedriver.GetDriver(cs)
	if len(_cDriver.Type) == 0 {
		panic("gcache> cache driver error")
	}
	// mq initialize
	switch _cDriver.Type {
	case "redis":
		cache := &gcacheredis.Redis{}
		cache.Driver = _cDriver
		ic = cache
	default:
		panic("gcache> unknown driver type")
	}
	return ic
}

// R : connect by region
func R(rs string) (ic gcachebase.ICache) {
	// init all
	gcacheregion.Init()
	// get db cluster
	_cRegion := gcacheregion.GetRegion(rs)
	if len(_cRegion.Name) == 0 {
		panic("gcache> cache region error")
	}
	// db initialize
	switch _cRegion.Driver.Type {
	case "redis":
		cache := &gcacheredis.Redis{}
		cache.Region = _cRegion
		cache.Driver = _cRegion.Driver
		ic = cache
	default:
		panic("gcache> unknown region driver type")
	}
	return ic
}
