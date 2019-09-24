package config

import (
	"base/util"
	"testing"
)

func TestConfigDatabase(t *testing.T) {
	Init()
	tmp := Load("database").GetStringMap("drivers")
	util.Dump(tmp)
	Init()
	all := Load("etcd").AllSettings()
	util.Dump(all)
}
