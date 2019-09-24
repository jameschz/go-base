package gcache

import (
	"base/gcache/base"
	"base/gcache/driver"
	"base/gcache/redis"
	"base/gcache/region"
)

// connect by driver
func D(cs string) (ic base.ICache) {
	// init driver
	driver.Init()
	// get mq driver
	c_driver := driver.GetDriver(cs)
	if len(c_driver.Type) == 0 {
		panic("gcache> cache driver error")
	}
	// mq initialize
	switch c_driver.Type {
	case "redis":
		cache := &redis.Redis{}
		cache.Driver = c_driver
		ic = cache
	default:
		panic("gcache> unknown driver type")
	}
	return ic
}

// connect by region
func R(rs string) (ic base.ICache) {
	// init all
	region.Init()
	// get db cluster
	c_region := region.GetRegion(rs)
	if len(c_region.Name) == 0 {
		panic("gcache> cache region error")
	}
	// db initialize
	switch c_region.Driver.Type {
	case "redis":
		cache := &redis.Redis{}
		cache.Region = c_region
		ic = cache
	default:
		panic("gcache> unknown region driver type")
	}
	return ic
}
