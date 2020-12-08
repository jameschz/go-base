package gcacheregion

import (
	"strconv"
	"strings"
	"time"

	"github.com/jameschz/go-base/lib/config"
	gcachedriver "github.com/jameschz/go-base/lib/gcache/driver"
)

// Region :
type Region struct {
	Driver *gcachedriver.Driver
	Name   string
	Desc   string
	TTL    int
}

var (
	_cRegions map[string]*Region
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

// Init :
func Init() bool {
	// init once
	if len(_cRegions) > 0 {
		return true
	}
	// init drivers
	gcachedriver.Init()
	// init regions
	regions := config.Load("cache").GetStringMap("regions")
	_cRegions = make(map[string]*Region, len(regions))
	for _rName, _rConf := range regions {
		// convert interface map
		_cRegion := _rConf.(map[string]interface{})
		// check driver
		region := &Region{
			Driver: gcachedriver.GetDriver(_cRegion["driver"].(string)),
			Name:   _cRegion["name"].(string),
			Desc:   _cRegion["desc"].(string),
			TTL:    stringToInt(_cRegion["ttl"].(string)),
		}
		// check region
		if len(region.Name) == 0 ||
			region.TTL <= 0 {
			panic("gcache> region format error")
		}
		// save region
		_cRegions[_rName] = region
	}
	return true
}

// GetRegion :
func GetRegion(rs string) (region *Region) {
	if _, r := _cRegions[rs]; !r {
		panic("gmq> can not find region")
	}
	region = _cRegions[rs]
	return region
}

// GetKey :
func (r *Region) GetKey(k string) string {
	return r.Name + "_" + k
}

// GetExp :
func (r *Region) GetExp() time.Duration {
	return time.Duration(r.TTL) * time.Second
}
