package region

import (
	"go-base/lib/config"
	"go-base/lib/gcache/driver"
	"strconv"
	"strings"
	"time"
)

type Region struct {
	Driver *driver.Driver
	Name   string
	Desc   string
	TTL    int
}

var (
	_c_regions map[string]*Region
)

// private static func
func stringToInt(secondStr string) int {
	digits := strings.Split(secondStr, "*")
	sum := 1
	for _, v := range digits {
		i, err := strconv.Atoi(strings.Trim(v, " "))
		if err != nil {
			panic("gcache> ttl format error")
		}
		sum *= i
	}
	return sum
}

// public static func
func Init() bool {
	// init once
	if len(_c_regions) > 0 {
		return true
	}
	// init drivers
	driver.Init()
	// init regions
	regions := config.Load("cache").GetStringMap("regions")
	_c_regions = make(map[string]*Region, len(regions))
	for r_name, r_conf := range regions {
		// convert interface map
		c_region := r_conf.(map[string]interface{})
		// check driver
		region := &Region{
			Driver: driver.GetDriver(c_region["driver"].(string)),
			Name:   c_region["name"].(string),
			Desc:   c_region["desc"].(string),
			TTL:    stringToInt(c_region["ttl"].(string)),
		}
		// check region
		if len(region.Name) == 0 ||
			region.TTL <= 0 {
			panic("gcache> region format error")
		}
		// save region
		_c_regions[r_name] = region
	}
	return true
}

// public static func
func GetRegion(rs string) (region *Region) {
	if _, r := _c_regions[rs]; !r {
		panic("gmq> can not find region")
	}
	region = _c_regions[rs]
	return region
}

// public member func
func (r *Region) GetKey(k string) string {
	return r.Name + ":" + k
}

// public member func
func (r *Region) GetExp() time.Duration {
	return time.Duration(r.TTL) * time.Second
}
