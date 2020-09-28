package gcachedriver

import (
	"github.com/jameschz/go-base/lib/config"
	"hash/crc32"
	"math/rand"
	"time"
)

// Driver :
type Driver struct {
	Type  string
	Algo  string
	Nodes []string
}

var (
	_cDrivers map[string]*Driver
)

// Init :
func Init() bool {
	if len(_cDrivers) > 0 {
		return true
	}
	drivers := config.Load("cache").GetStringMap("drivers")
	_cDrivers = make(map[string]*Driver, len(drivers))
	for _cName, _cConf := range drivers {
		// convert interface map
		_cDriver := _cConf.(map[string]interface{})
		// get nodes config
		_cNodesList := _cDriver["nodes"].([]interface{})
		_cNodes := make([]string, len(_cNodesList))
		for k, v := range _cNodesList {
			_cNodes[k] = v.(string)
		}
		// check driver
		driver := &Driver{
			Type:  _cDriver["type"].(string),
			Algo:  _cDriver["algo"].(string),
			Nodes: _cNodes,
		}
		// check driver
		if len(driver.Algo) == 0 ||
			len(driver.Nodes) == 0 {
			panic("gmq> init driver error")
		}
		// save driver
		_cDrivers[_cName] = driver
	}
	return true
}

// GetDriver :
func GetDriver(ds string) (driver *Driver) {
	if _, r := _cDrivers[ds]; !r {
		panic("gmq> can not find driver")
	}
	driver = _cDrivers[ds]
	return driver
}

// GetShardNode :
func (d *Driver) GetShardNode(s string) string {
	var sk int
	switch d.Algo {
	case "rand":
		rand.Seed(time.Now().UnixNano())
		sk = rand.Intn(len(d.Nodes))
	case "hash":
		code := crc32.ChecksumIEEE([]byte(s))
		sk = int(code) % len(d.Nodes)
	}
	return d.Nodes[sk]
}
