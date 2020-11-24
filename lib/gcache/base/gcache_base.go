package gcachebase

import (
	"time"

	"github.com/go-redis/redis"
	gcachedriver "github.com/jameschz/go-base/lib/gcache/driver"
	gcacheregion "github.com/jameschz/go-base/lib/gcache/region"
)

// DataSource :
type DataSource struct {
	ID        string
	Name      string
	Node      string        // connection node
	RedisConn *redis.Client // redis connection
}

// Cache :
type Cache struct {
	Node       string               // connection node
	Driver     *gcachedriver.Driver // driver ptr
	Region     *gcacheregion.Region // region ptr
	DataSource *DataSource          // datasource
	RedisConn  *redis.Client        // redis connection
}

// ICache :
type ICache interface {
	Connect(k string) error
	Close() error
	Set(k string, v string) error
	SetTTL(k string, v string, exp time.Duration) error
	Get(k string) (string, error)
	Del(k string) error
	Incr(k string) (int64, error)
	IncrBy(k string, val int64) (int64, error)
}
