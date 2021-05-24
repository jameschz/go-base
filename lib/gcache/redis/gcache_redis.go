package gcacheredis

import (
	"time"

	"github.com/go-redis/redis"
	gcachebase "github.com/jameschz/go-base/lib/gcache/base"
	gcachepool "github.com/jameschz/go-base/lib/gcache/pool"
)

var (
	_cTimeout = time.Hour
)

// Redis :
type Redis struct {
	gcachebase.Cache // extends base cache
}

// Connect :
func (r *Redis) Connect(k string) error {
	// connect once
	if r.RedisConn == nil {
		// gdopool
		gcachepool.Init()
		name := r.Driver.Name
		node := r.Driver.GetShardNode(k)
		dataSource, err := gcachepool.Fetch(name, node)
		// init redis vars
		r.Node = node
		r.DataSource = dataSource
		r.RedisConn = dataSource.RedisConn
		return err
	}
	return nil
}

// Close :
func (r *Redis) Close() error {
	// close once
	if r.RedisConn != nil {
		gcachepool.Return(r.DataSource)
		r.DataSource = nil
		r.RedisConn = nil
		r.Node = ""
		return nil
	}
	return nil
}

// Set :
func (r *Redis) Set(k string, v string) error {
	// default exp
	exp := _cTimeout
	// use region exp
	if r.Region != nil {
		exp = r.Region.GetExp()
	}
	return r.SetTTL(k, v, exp)
}

// SetTTL :
func (r *Redis) SetTTL(k string, v string, exp time.Duration) error {
	// use region key
	if r.Region != nil {
		k = r.Region.GetKey(k)
	}
	// connect
	r.Connect(k)
	// set kv
	_, err := r.RedisConn.Set(k, v, exp).Result()
	if err != nil {
		return err
	}
	return nil
}

// Get :
func (r *Redis) Get(k string) (string, error) {
	// use region key
	if r.Region != nil {
		k = r.Region.GetKey(k)
	}
	// connect
	r.Connect(k)
	// get kv
	v, err := r.RedisConn.Get(k).Result()
	if err == redis.Nil {
		return "", nil
	} else if err != nil {
		return "", err
	}
	return v, nil
}

// Del :
func (r *Redis) Del(k string) error {
	// use region key
	if r.Region != nil {
		k = r.Region.GetKey(k)
	}
	// connect
	r.Connect(k)
	// del kv
	_, err := r.RedisConn.Del(k).Result()
	if err == redis.Nil {
		return nil
	} else if err != nil {
		return err
	}
	return nil
}

// Incr :
func (r *Redis) Incr(k string) (int64, error) {
	// use region key
	if r.Region != nil {
		k = r.Region.GetKey(k)
	}
	// connect
	r.Connect(k)
	// get kv
	v, err := r.RedisConn.Incr(k).Result()
	if err == redis.Nil {
		return 0, nil
	} else if err != nil {
		return 0, err
	}
	return v, nil
}

// IncrBy :
func (r *Redis) IncrBy(k string, val int64) (int64, error) {
	// use region key
	if r.Region != nil {
		k = r.Region.GetKey(k)
	}
	// connect
	r.Connect(k)
	// get kv
	v, err := r.RedisConn.IncrBy(k, val).Result()
	if err == redis.Nil {
		return 0, nil
	} else if err != nil {
		return 0, err
	}
	return v, nil
}
