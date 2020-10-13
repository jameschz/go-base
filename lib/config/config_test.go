package config

import (
	"testing"

	"github.com/jameschz/go-base/lib/gutil"
)

func TestConfigDatabase(t *testing.T) {
	Init()
	tmp := Load("database").GetStringMap("drivers")
	gutil.Dump(tmp)
	Init()
	all := Load("etcd").AllSettings()
	gutil.Dump(all)
}
