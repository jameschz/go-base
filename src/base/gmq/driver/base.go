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
	_mq_drivers map[string]*Driver
)

func Init() bool {
	if len(_mq_drivers) > 0 {
		return true
	}
	drivers := config.Load("queue").GetStringMap("drivers")
	_mq_drivers = make(map[string]*Driver, len(drivers))
	for mq_name, mq_conf := range drivers {
		// convert interface map
		mq_driver := mq_conf.(map[string]interface{})
		// get nodes config
		nodes_a := mq_driver["nodes"].([]interface{})
		mq_nodes := make([]string, len(nodes_a))
		for k, v := range nodes_a {
			mq_nodes[k] = v.(string)
		}
		// check driver
		driver := &Driver{
			Type:  mq_driver["type"].(string),
			Algo:  mq_driver["algo"].(string),
			Nodes: mq_nodes,
		}
		// check driver
		if len(driver.Algo) == 0 ||
			len(driver.Nodes) == 0 {
			panic("gmq> init driver error")
		}
		// save driver
		_mq_drivers[mq_name] = driver
	}
	return true
}

func GetDriver(ds string) (driver *Driver) {
	if _, r := _mq_drivers[ds]; !r {
		panic("gmq> can not find driver")
	}
	driver = _mq_drivers[ds]
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
