package base

import (
	"base/gcache/driver"
	"base/gcache/region"
	"time"
)

type Cache struct {
	Driver   *driver.Driver // driver ptr
	Region   *region.Region // region ptr
	NodeName string         // node name
}

type ICache interface {
	Connect(node string) error
	Close() error
	Shard(k string) (err error)
	Set(k string, v string) error
	SetTTL(k string, v string, exp time.Duration) error
	Get(k string) (string, error)
	Del(k string) error
}
