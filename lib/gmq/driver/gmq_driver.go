package gmqdriver

import (
	"hash/crc32"
	"math/rand"
	"time"

	"github.com/jameschz/go-base/lib/config"
)

// Driver :
type Driver struct {
	Type  string
	Algo  string
	Nodes []string
}

var (
	_mqDrivers map[string]*Driver
)

// Init :
func Init() bool {
	if len(_mqDrivers) > 0 {
		return true
	}
	drivers := config.Load("queue").GetStringMap("drivers")
	_mqDrivers = make(map[string]*Driver, len(drivers))
	for _mqName, _mqConf := range drivers {
		// convert interface map
		_mqDriver := _mqConf.(map[string]interface{})
		// get nodes config
		_mqNodesList := _mqDriver["nodes"].([]interface{})
		_mqNodes := make([]string, len(_mqNodesList))
		for k, v := range _mqNodesList {
			_mqNodes[k] = v.(string)
		}
		// check driver
		driver := &Driver{
			Type:  _mqDriver["type"].(string),
			Algo:  _mqDriver["algo"].(string),
			Nodes: _mqNodes,
		}
		// check driver
		if len(driver.Algo) == 0 ||
			len(driver.Nodes) == 0 {
			panic("gmq> init driver error")
		}
		// save driver
		_mqDrivers[_mqName] = driver
	}
	return true
}

// GetDriver :
func GetDriver(ds string) (driver *Driver) {
	if _, r := _mqDrivers[ds]; !r {
		panic("gmq> can not find driver")
	}
	driver = _mqDrivers[ds]
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
