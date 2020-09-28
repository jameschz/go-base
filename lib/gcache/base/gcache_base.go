package gcachebase

import (
	"github.com/jameschz/go-base/lib/gcache/driver"
	"github.com/jameschz/go-base/lib/gcache/region"
	"time"
)

// Cache :
type Cache struct {
	Driver   *gcachedriver.Driver // driver ptr
	Region   *gcacheregion.Region // region ptr
	NodeName string               // node name
}

// ICache :
type ICache interface {
	Connect(node string) error
	Close() error
	Shard(k string) (err error)
	Set(k string, v string) error
	SetTTL(k string, v string, exp time.Duration) error
	Get(k string) (string, error)
	Del(k string) error
}
