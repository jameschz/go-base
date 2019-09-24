package driver

import (
	"base/config"
	"hash/crc32"
	"math/rand"
	"time"
)

type Driver struct {
	Type  string
	Algo  string
	Nodes []string
}

var (
	_c_drivers map[string]*Driver
)

func Init() bool {
	if len(_c_drivers) > 0 {
		return true
	}
	drivers := config.Load("cache").GetStringMap("drivers")
	_c_drivers = make(map[string]*Driver, len(drivers))
	for c_name, c_conf := range drivers {
		// convert interface map
		c_driver := c_conf.(map[string]interface{})
		// get nodes config
		nodes_a := c_driver["nodes"].([]interface{})
		c_nodes := make([]string, len(nodes_a))
		for k, v := range nodes_a {
			c_nodes[k] = v.(string)
		}
		// check driver
		driver := &Driver{
			Type:  c_driver["type"].(string),
			Algo:  c_driver["algo"].(string),
			Nodes: c_nodes,
		}
		// check driver
		if len(driver.Algo) == 0 ||
			len(driver.Nodes) == 0 {
			panic("gmq> init driver error")
		}
		// save driver
		_c_drivers[c_name] = driver
	}
	return true
}

func GetDriver(ds string) (driver *Driver) {
	if _, r := _c_drivers[ds]; !r {
		panic("gmq> can not find driver")
	}
	driver = _c_drivers[ds]
	return driver
}

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
